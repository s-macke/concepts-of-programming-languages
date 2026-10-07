import assert from 'node:assert/strict';

const delay = ms => new Promise(resolve => setTimeout(resolve, ms));
function deferred() {
  let resolve;
  const promise = new Promise(done => { resolve = done; });
  return { promise, resolve };
}

// Called with a fresh catalog room. All intercepted requests target the test
// server; this never interacts with an existing lecture or development room.
export async function checkGroupRaces(browser, host) {
  const link = await host.getByRole('textbox', { name: 'Join link' }).inputValue();
  const id = link.split('/').at(-1);
  const context = await browser.newContext();
  const participant = await context.newPage();
  let joins = 0;
  let stateRequests = 0;
  const pendingState = deferred();
  const releaseState = deferred();
  await participant.route(`**/api/groups/${id}/join`, async route => {
    joins++;
    if (joins === 1) await route.fulfill({ status: 503, json: { error: 'Join temporarily unavailable' } });
    else await route.continue();
  });
  await participant.route(`**/api/groups/${id}`, async route => {
    stateRequests++;
    if (stateRequests === 1) {
      pendingState.resolve();
      await releaseState.promise;
      await route.fulfill({ status: 503, json: { error: 'State temporarily unavailable' } });
    } else await route.continue();
  });
  try {
    await participant.goto(link);
    const join = participant.getByRole('button', { name: 'Join', exact: true });
    await join.click();
    await participant.getByRole('status').filter({ hasText: 'Join temporarily unavailable' }).waitFor();
    // A rejected join must still be retryable.
    const originalButton = await join.elementHandle();
    await join.click();
    await pendingState.promise;
    assert.equal(await join.count(), 0, 'successful join removes Join before state arrives');
    const credential = await participant.evaluate(id => sessionStorage.getItem(`quiz-group:${id}`), id);
    assert.ok(credential);
    // Even a retained reference must not create a second participant/poll loop.
    await originalButton.evaluate(button => button.click());
    await delay(100);
    assert.equal(joins, 2, 'one failed join plus exactly one successful join');
    releaseState.resolve();
    await participant.getByRole('status').filter({ hasText: 'State temporarily unavailable' }).waitFor();
    assert.equal(await join.count(), 0, 'failed state fetch cannot offer rejoining');
    await participant.getByRole('heading', { name: "You're in!", exact: true }).waitFor();
    assert.equal(await participant.evaluate(id => sessionStorage.getItem(`quiz-group:${id}`), id), credential, 'retry uses the same identity');
    await host.getByText('1 participants', { exact: true }).waitFor();
    assert.equal(joins, 2);
    assert.equal(stateRequests, 2, 'only one poll loop retries the failed state request');
  } finally {
    releaseState.resolve();
    await context.close();
  }

  // Both successful and failed stale list responses must leave the group view
  // mounted. Navigate back/forward while each list request is still pending.
  for (const status of [200, 503]) {
    const pendingList = deferred();
    const releaseList = deferred();
    const handler = async route => {
      pendingList.resolve();
      await releaseList.promise;
      await route.fulfill({ status, json: status === 200 ? [] : { error: 'Stale list error' } });
    };
    await host.route('**/api/quizzes', handler);
    try {
      await host.goBack();
      await pendingList.promise;
      await host.goForward();
      await host.getByRole('heading', { name: 'Scan to join' }).waitFor();
      const responsePromise = host.waitForResponse('**/api/quizzes');
      releaseList.resolve();
      await (await responsePromise).finished();
      await delay(100); // Allow response.json() and the render microtasks to finish.
      assert.equal(await host.getByRole('heading', { name: 'Scan to join' }).count(), 1, `stale ${status} list response cannot replace group`);
      assert.equal(await host.getByRole('button', { name: 'Start quiz' }).count(), 1);
      assert.equal(new URL(host.url()).hash, new URL(link).hash);
    } finally {
      releaseList.resolve();
      await host.unroute('**/api/quizzes', handler);
    }
  }
  console.log('PASS: delayed/failed initial state cannot duplicate a join; stale successful/failed quiz-list responses cannot overwrite group navigation.');
}

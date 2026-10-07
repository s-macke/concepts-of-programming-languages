// Run after npm run build. Install a Playwright Chromium browser, or set
// QUIZ_BROWSER_BIN to a local Chrome/Chromium executable.
import { chromium } from 'playwright';
import jsQR from 'jsqr';
import { checkGroupRaces } from './group-races.mjs';
import assert from 'node:assert/strict';
import { spawn, execFileSync } from 'node:child_process';
import { mkdtempSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import { createServer } from 'node:net';

const root = fileURLToPath(new URL('../../../../', import.meta.url));
const temp = mkdtempSync(join(tmpdir(), 'quiz-group-test-'));
const probe = createServer();
await new Promise(resolve => probe.listen(0, '127.0.0.1', resolve));
const port = probe.address().port;
await new Promise(resolve => probe.close(resolve));
const origin = `http://127.0.0.1:${port}`;
let server, browser;
const errors = [];
const fixture = `title: Browser group test
questions:
- text: Pick A
  options: [{id: a, text: A}, {id: b, text: B}]
  answer: a
  explanation: A is correct.
- type: multiple
  text: Pick both
  options: [{id: a, text: A}, {id: b, text: B}]
  answer: [a, b]
- type: text
  text: Keyword
  answer: go
- type: truefalse
  text: Is Go a programming language?
  answer: "true"
`;
const delay = ms => new Promise(resolve => setTimeout(resolve, ms));
async function contains(page, text) {
  await page.getByText(text, { exact: true }).waitFor({ timeout: 12000 });
}
async function page(context) {
  const p = await context.newPage();
  p.on('pageerror', e => errors.push(e.message));
  return p;
}
try {
  execFileSync('go', ['build', '-o', join(temp, 'quizserver'), './src/quiz/cmd/quizserver'], { cwd: root, stdio: 'pipe' });
  server = spawn(join(temp, 'quizserver'), ['-addr', `127.0.0.1:${port}`, '-public-url', origin], { cwd: root, stdio: 'ignore' });
  for (let i = 0; i < 80; i++) {
    try { if ((await fetch(origin + '/api/quizzes')).ok) break; } catch {}
    if (i === 79) throw new Error('Quiz server did not start');
    await delay(100);
  }
  browser = await chromium.launch({ headless: true, executablePath: process.env.QUIZ_BROWSER_BIN || undefined, args: ['--no-sandbox'] });
  const host = await page(await browser.newContext({ viewport: { width: 1280, height: 900 } }));
  const phone = await page(await browser.newContext({ viewport: { width: 390, height: 844 } }));
  const second = await page(await browser.newContext());
  const late = await page(await browser.newContext());
  await host.goto(origin);
  await host.getByRole('button', { name: 'Host group quiz', exact: true }).first().click();
  await host.getByRole('heading', { name: 'Scan to join' }).waitFor();
  await checkGroupRaces(browser, host);
  await host.getByRole('button', { name: 'Delete room' }).click();
  await host.getByRole('button', { name: 'Back to quizzes' }).click();
  await host.getByLabel('Host uploaded quiz as a group').check();
  await host.locator('#upload').setInputFiles({ name: 'group.yaml', mimeType: 'application/yaml', buffer: Buffer.from(fixture) });
  await host.getByRole('heading', { name: 'Scan to join' }).waitFor();
  await host.screenshot({ path: join(temp, 'host-lobby.png') });
  const link = await host.getByRole('textbox', { name: 'Join link' }).inputValue();
  const pixels = await host.locator('canvas').evaluate(canvas => ({ width: canvas.width, height: canvas.height, data: Array.from(canvas.getContext('2d').getImageData(0, 0, canvas.width, canvas.height).data) }));
  assert.equal(jsQR(new Uint8ClampedArray(pixels.data), pixels.width, pixels.height)?.data, link, 'QR decodes to the join link');
  for (const p of [phone, second]) { await p.goto(link); await p.getByRole('button', { name: 'Join', exact: true }).click(); await contains(p, "You're in!"); }
  await contains(host, '2 participants');
  await host.reload(); await host.getByRole('button', { name: 'Start quiz' }).waitFor();
  await host.getByRole('button', { name: 'Start quiz' }).click();
  await phone.getByRole('radio', { name: 'A', exact: true }).check();
  await delay(2300);
  assert.equal(await phone.getByRole('radio', { name: 'A', exact: true }).isChecked(), true, 'poll preserves selected input');
  await phone.getByRole('button', { name: 'Submit answer' }).click();
  await contains(phone, 'Answer submitted. Waiting for the host to reveal results…');
  await phone.reload(); await contains(phone, 'Answer submitted. Waiting for the host to reveal results…');
  await second.getByRole('radio', { name: 'B', exact: true }).check();
  await second.getByRole('button', { name: 'Submit answer' }).click();
  await contains(host, '2 participants · 2 submitted');
  await host.getByRole('button', { name: 'Reveal results' }).click();
  await contains(phone, 'Your answer was correct (+1).');
  await contains(second, 'Your answer was not correct.');
  await phone.screenshot({ path: join(temp, 'mobile-reveal.png') });
  await contains(host, '2 submitted · 0 unanswered · 1 correct · 1 incorrect');
  await late.goto(link); await late.getByRole('button', { name: 'Join', exact: true }).click();
  await contains(late, 'You did not answer this question.');
  await contains(host, '2 submitted · 1 unanswered · 1 correct · 1 incorrect');
  await host.getByRole('button', { name: 'Next question' }).click();
  for (const name of ['A', 'B']) await phone.getByRole('checkbox', { name, exact: true }).check();
  await phone.getByRole('button', { name: 'Submit answer' }).click();
  await contains(host, '3 participants · 1 submitted');
  await host.getByRole('button', { name: 'Reveal results' }).click();
  await contains(phone, 'A: 1 (100%)'); await contains(phone, 'B: 1 (100%)');
  await host.getByRole('button', { name: 'Next question' }).click();
  await phone.getByRole('textbox').fill(' GO ');
  await phone.context().setOffline(true); await delay(2500);
  await contains(phone, 'Connection interrupted. Retrying… Failed to fetch');
  await phone.context().setOffline(false); await delay(2500);
  assert.equal(await phone.getByRole('textbox').inputValue(), ' GO ', 'reconnection preserves typed answer');
  await phone.getByRole('button', { name: 'Submit answer' }).click();
  await second.getByRole('textbox').fill('A private name');
  await second.getByRole('button', { name: 'Submit answer' }).click();
  await contains(host, '3 participants · 2 submitted');
  await host.getByRole('button', { name: 'Reveal results' }).click();
  await contains(phone, 'Expected: go');
  assert.equal((await host.locator('body').innerText()).includes('A private name'), false);
  await host.getByRole('button', { name: 'Next question' }).click();
  await phone.getByRole('radio', { name: 'True', exact: true }).check();
  await phone.getByRole('button', { name: 'Submit answer' }).click();
  await contains(host, '3 participants · 1 submitted');
  await host.getByRole('button', { name: 'Reveal results' }).click();
  await host.getByRole('button', { name: 'Finish quiz', exact: true }).click();
  await contains(host, '50.0% average score'); await contains(phone, '50.0% average score');
  await contains(host, '1 completed the entire quiz · 6 total submissions');
  assert.equal(await phone.evaluate(() => document.documentElement.scrollWidth <= innerWidth), true, 'mobile has no horizontal overflow');
  await phone.screenshot({ path: join(temp, 'mobile-summary.png') });
  await host.screenshot({ path: join(temp, 'host-summary.png') });
  await host.getByRole('button', { name: 'Delete room' }).click();
  await contains(phone, 'unknown or expired room');
  assert.deepEqual(errors, []);
  console.log('PASS: QR decoding, hosted catalog/upload, all question types, independent participants, late join, aggregate results, refresh, reconnect, input preservation, mobile layout, deletion.');
  if (process.env.QUIZ_SCREENSHOT_DIR) {
    const { copyFileSync, mkdirSync } = await import('node:fs');
    mkdirSync(process.env.QUIZ_SCREENSHOT_DIR, { recursive: true });
    for (const name of ['mobile-summary.png', 'host-summary.png', 'host-lobby.png', 'mobile-reveal.png']) copyFileSync(join(temp, name), join(process.env.QUIZ_SCREENSHOT_DIR, name));
  }
} finally {
  await browser?.close();
  server?.kill('SIGINT');
  if (server && server.exitCode === null) await new Promise(resolve => server.once('exit', resolve));
  rmSync(temp, { recursive: true, force: true });
}

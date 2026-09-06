// Copyright 2026 The Concepts of Programming Languages Authors.
// Licensed under the Apache License, Version 2.0.

(function() {
  'use strict';

  var playgroundEndpoint = 'https://play.golang.org/compile';

  // Extract the original program from the highlighted code DOM. Hidden
  // prefix and suffix blocks contain the parts omitted from the slide.
  function codeText(node) {
    var source = '';
    for (var i = 0; i < node.childNodes.length; i++) {
      var child = node.childNodes[i];
      if (child.nodeType === Node.TEXT_NODE) {
        source += child.nodeValue;
        continue;
      }
      if (child.nodeType !== Node.ELEMENT_NODE) continue;
      if (child.tagName === 'BUTTON') continue;
      if (child.tagName === 'SPAN' && child.className === 'number') continue;
      if (
        child.tagName === 'DIV' ||
        child.tagName === 'BR' ||
        child.tagName === 'PRE'
      ) {
        source += '\n';
      }
      source += codeText(child);
    }
    return source.replace(/\u00a0/g, ' ');
  }

  function appendOutput(element, kind, message) {
    if (!message) return;

    // A form feed clears the Playground terminal.
    var chunks = message.split('\x0c');
    if (chunks.length > 1) {
      element.textContent = '';
      message = chunks[chunks.length - 1];
    }

    var shouldScroll =
      element.scrollTop + element.offsetHeight >= element.scrollHeight;
    var span = document.createElement('span');
    span.className = kind;
    span.textContent = message;
    element.appendChild(span);
    if (shouldScroll) {
      element.scrollTop = element.scrollHeight - element.offsetHeight;
    }
  }

  function initPlayground(code) {
    var output = document.createElement('div');
    var outputText = document.createElement('pre');
    var request;
    var playbackTimer;
    var sequence = 0;

    function stop(showMessage) {
      sequence++;
      if (request) {
        request.abort();
        request = null;
      }
      if (playbackTimer) {
        window.clearTimeout(playbackTimer);
        playbackTimer = null;
      }
      if (showMessage) {
        appendOutput(outputText, 'system', '\nProgram exited: killed.');
      }
    }

    function playEvents(data, currentSequence) {
      var events = (data.Events || []).slice();

      function next() {
        if (currentSequence !== sequence) return;
        if (events.length === 0) {
          if (data.IsTest) {
            if (data.TestsFailed > 0) {
              appendOutput(
                outputText,
                'system',
                '\n' + data.TestsFailed + ' test' +
                  (data.TestsFailed === 1 ? '' : 's') + ' failed.'
              );
            } else {
              appendOutput(outputText, 'system', '\nAll tests passed.');
            }
          } else if (data.Status > 0) {
            appendOutput(
              outputText,
              'system',
              '\nProgram exited: status ' + data.Status + '.'
            );
          } else if (data.Errors) {
            appendOutput(outputText, 'system', '\n' + data.Errors + '.');
          }
          return;
        }

        var event = events.shift();
        var delay = Math.max(0, event.Delay || 0) / 1000000;
        if (delay === 0) {
          appendOutput(outputText, event.Kind, event.Message);
          next();
          return;
        }
        playbackTimer = window.setTimeout(function() {
          playbackTimer = null;
          appendOutput(outputText, event.Kind, event.Message);
          next();
        }, delay);
      }

      next();
    }

    function onRun() {
      stop(false);
      var currentSequence = sequence;
      output.style.display = 'block';
      outputText.textContent = '';
      primaryRun.style.display = 'none';

      var form = new URLSearchParams();
      form.set('version', '2');
      form.set('body', codeText(code));
      form.set('withVet', 'true');

      request = new AbortController();
      fetch(playgroundEndpoint, {
        method: 'POST',
        body: form,
        signal: request.signal,
      })
        .then(function(response) {
          if (!response.ok) {
            throw new Error('Playground returned HTTP ' + response.status + '.');
          }
          return response.json();
        })
        .then(function(data) {
          if (currentSequence !== sequence) return;
          request = null;

          if (data.Errors && data.Errors !== 'process took too long') {
            appendOutput(outputText, 'stderr', data.Errors);
            appendOutput(outputText, 'system', '\nGo build failed.');
            return;
          }
          if (data.VetErrors) {
            appendOutput(outputText, 'stderr', data.VetErrors);
            appendOutput(outputText, 'system', '\nGo vet exited.\n\n');
          }
          playEvents(data, currentSequence);
        })
        .catch(function(error) {
          if (error.name === 'AbortError' || currentSequence !== sequence) return;
          request = null;
          appendOutput(
            outputText,
            'stderr',
            'Error communicating with the Go Playground: ' + error.message
          );
        });
    }

    function onStop() {
      stop(true);
    }

    function onClose() {
      stop(false);
      output.style.display = 'none';
      primaryRun.style.display = 'inline-block';
    }

    var primaryRun = document.createElement('button');
    primaryRun.type = 'button';
    primaryRun.className = 'run';
    primaryRun.textContent = 'Run';
    primaryRun.addEventListener('click', onRun, false);

    var outputRun = document.createElement('button');
    outputRun.type = 'button';
    outputRun.className = 'run';
    outputRun.textContent = 'Run';
    outputRun.addEventListener('click', onRun, false);

    var stopButton = document.createElement('button');
    stopButton.type = 'button';
    stopButton.className = 'kill';
    stopButton.textContent = 'Stop';
    stopButton.addEventListener('click', onStop, false);

    var closeButton = document.createElement('button');
    closeButton.type = 'button';
    closeButton.className = 'close';
    closeButton.textContent = 'Close';
    closeButton.addEventListener('click', onClose, false);

    var primaryButtons = document.createElement('div');
    primaryButtons.className = 'buttons';
    primaryButtons.appendChild(primaryRun);
    code.parentNode.insertBefore(primaryButtons, code.nextSibling);

    var outputButtons = document.createElement('div');
    outputButtons.className = 'buttons';
    outputButtons.appendChild(outputRun);
    outputButtons.appendChild(stopButton);
    outputButtons.appendChild(closeButton);

    output.className = 'output';
    output.style.display = 'none';
    output.appendChild(outputButtons);
    output.appendChild(outputText);
    primaryButtons.parentNode.insertBefore(output, primaryButtons.nextSibling);
  }

  var playgrounds = document.querySelectorAll('div.playground');
  for (var i = 0; i < playgrounds.length; i++) {
    initPlayground(playgrounds[i]);
  }
})();

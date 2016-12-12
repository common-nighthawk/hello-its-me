package templates

import(
  "../secrets"
  "fmt"
)

var gaScript = fmt.Sprintf(`(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
                            (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
                            m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
                            })(window,document,'script','https://www.google-analytics.com/analytics.js','ga');
                            ga('create', '%s', 'auto');
                            ga('send', 'pageview');`, secrets.GACode)

func MsgScript() string {
  return `
    var start = document.getElementById('start');
    var stop = document.getElementById('stop');
    var dismiss = document.getElementById('dismiss');
    var send = document.getElementById('send');
    var message = document.getElementById('message')
    var error = document.getElementsByClassName("error")[0];
    var rec = document.getElementById('rec');
    setInitialStyle();

    var media = { tag: 'audio', type: 'audio/webm', gUM: {audio: true} };
    var recorder, chunks, blob;

    if (navigator.mediaDevices) {
      navigator.mediaDevices.getUserMedia(media.gUM)
      .then(function(stream) {
        recorder = new MediaRecorder(stream);
        recorder.ondataavailable = e => { chunks.push(e.data); };
        console.log('got media successfully');
      }).catch(function(err) { displayGetUserMediaError(err); });
    } else {
      displayGetUserMediaNoGo();
    }

    start.onclick = e => {
      chunks=[];
      recorder.start();
      start.style.display='none'; stop.style.display='initial'; rec.style.display='block';
    }

    stop.onclick = e => {
      recorder.stop();
      stop.style.display='none'; rec.style.display='none'; send.style.display='initial'; dismiss.style.display='initial';
      setBlob();
    }

    send.onclick = e => {
      var formData = new FormData();
      formData.append('blob', blob);
      postMessage(formData);
      dismiss.style.display='none'; send.style.display='none';
    }

    dismiss.onclick = e => {
      window.location.reload(false);
    }

    function setBlob() {
      if (chunks.length < 1) {
        setTimeout(setBlob, 25);
        return;
      }

      blob = new Blob(chunks, {type: media.type });
      var url = window.URL.createObjectURL(blob);
      console.log('created audio successfully: ' + url);

      var ah = document.getElementById('audio-holder');
      var pa = document.createElement('p');
      var mt = document.createElement(media.tag);
      var hf = document.createElement('a');
      mt.controls = true;
      mt.src = url;
      pa.appendChild(mt);
      ah.appendChild(pa);
      message.innerHTML = 'You can preview your message. Then decide what you wanna do.'
    }

    function postMessage(formData) {
      var xhr = new XMLHttpRequest();
      var postURL = '/message?receiver_username=' + send.value + '&explode=' + start.value
      xhr.open('POST', postURL, true);
      xhr.send(formData);
      xhr.onload = function() {
        if (xhr.status == 200) {
          message.innerHTML = 'Your message was successfully sent!'
        } else {
          message.innerHTML = ''
          error.innerHTML = 'Oops. There was an error sending your message. Bleh--I\'m sorry :('
        }
      }
    }

    function setInitialStyle() {
      error.innerHTML = '';
      message.style.display='block';
      start.style.display='initial';
      stop.style.display='none';
      dismiss.style.display='none';
      send.style.display='none';
      rec.style.display='none';
    }

    function displayGetUserMediaError(err) {
      var defaultError =     'Uh oh. There was an error accessing the microphone.'
      var permissiondError = 'Uh oh. Access to your microphone is currently not allowed--\
                              making it impossible to send a message.<br>\
                              You can grant access with these instructions for\
                              <a href="https://support.google.com/chrome/answer/2693767">Chrome</a> and\
                              <a href="https://support.mozilla.org/en-US/kb/permissions-manager-give-ability-store-passwords-set-cookies-more">Firefox</a>.'
      var notFoundError =    'Uh oh. No microphone can be found--\
                              which is necessary to send a message.\
                              Go plug in a pair of headphones with a mic!'

      var errorMessage = '';
      switch(err["name"]) {
        case 'PermissionDeniedError':
          errorMessage = permissiondError
          break;
        case 'NotFoundError':
          errorMessage = notFoundError
          break;
        default:
          errorMessage = defaultError
      }

      start.style.display='none'
      error.innerHTML = errorMessage
      message.innerHTML = 'Reload the page once the mic is accessible.'
      console.log(err)
    }

    function displayGetUserMediaNoGo() {
      error.innerHTML = 'Bad news. Your browser does not support access to your microphone,\
                         which is the only way to send a message.\
                         Try using Chrome or Firefox.<br>\
                         You can still listen to messages with your current browser.'
      start.style.display='none';
      message.style.display='none';
    }
    `
}

func ExpireScript() string {
  return `
    var audioPlayers = document.getElementsByTagName("audio")
    for (var i = 0; i < audioPlayers.length; i++) {
      audioPlayers[i].addEventListener('play', function() {
        setExpiresAt(this.id);
      }, false);
    }

  function setExpiresAt(file) {
    var xhr = new XMLHttpRequest();
    var postURL = '/message_update?file=' + file;
    xhr.open('POST', postURL, true);
    xhr.send();
  }
  `
}

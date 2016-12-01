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
    var rec = document.getElementById('rec');

    stop.style.display='none';
    dismiss.style.display='none';
    send.style.display='none';
    rec.style.display='none';

    var media = { tag: 'audio', type: 'audio/mp3', gUM: {audio: true} };
    var stream, recorder, chunks, blob;

    navigator.mediaDevices.getUserMedia(media.gUM).then(_stream => {
      stream = _stream;
      recorder = new MediaRecorder(stream);
      recorder.ondataavailable = e => {
        chunks.push(e.data);
      };
      console.log('got media successfully');
    });

    start.onclick = e => {
      chunks=[];
      recorder.start();
      start.style.display='none'; stop.style.display='initial'; rec.style.display='block';
    }

    stop.onclick = e => {
      recorder.stop();
      stop.style.display='none'; rec.style.display='none'; send.style.display='initial'; dismiss.style.display='initial';

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

    send.onclick = e => {
      var formData = new FormData();
      formData.append('blob', blob);

      var xhr = new XMLHttpRequest();
      var postURL = '/message?receiver_username=' + stop.value
      xhr.open('POST', postURL, true);
      xhr.send(formData);

      dismiss.style.display='none'; send.style.display='none';
      message.innerHTML = 'Your message was successfully sent!'
    }

    dismiss.onclick = e => {
      window.location.reload(false);
    }`
}

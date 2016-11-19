package templates

func Script() string {
  return `
    var start = document.getElementById('start');
    var stop = document.getElementById('stop');
    var body = document.body;
    var media = { tag: 'audio', type: 'audio/ogg', ext: '.ogg', gUM: {audio: true} };
    var stream;
    var recorder;
    var chunks;
    var blob;
    var url;

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
    }

    stop.onclick = e => {
      recorder.stop();
      blob = new Blob(chunks, {type: media.type });
      url = window.URL.createObjectURL(blob);
      console.log('created audio successfully: ' + url);

      var pa = document.createElement('p');
      var mt = document.createElement(media.tag);
      var hf = document.createElement('a');
      mt.controls = true;
      mt.src = url;
      pa.appendChild(mt);
      body.appendChild(pa);

      var formData = new FormData();
      formData.append('key', blob);

      var xhr = new XMLHttpRequest();
      xhr.open('POST', 'savefile', true);
      xhr.onload = function(e) {};
      xhr.send(formData);
    }`
}

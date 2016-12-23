var audioPlayers = document.getElementsByTagName("audio")

for (var i = 0; i < audioPlayers.length; i++) {
  audioPlayers[i].addEventListener('play', function() {
    setExpiresAt(this.id);
  }, false);
}

function setExpiresAt(file) {
  var xhr = new XMLHttpRequest();
  var postURL = '/message_update?expire=true&file=' + file;
  xhr.open('POST', postURL, true);
  xhr.send();
}

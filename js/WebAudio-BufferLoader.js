// http://middleearmedia.com/demos/webaudio/bufferloader.html
// http://middleearmedia.com/web-audio-api-bufferloader/

function BufferLoader(context, urlList, callback) {
	this.context = context;
    this.urlList = urlList;
    this.onload = callback;
    this.bufferList = new Array();
    this.loadCount = 0;
}

BufferLoader.prototype.loadBuffer = function(url, index) {
    var request = new XMLHttpRequest();
    request.open("GET", url, true);
    request.responseType = "arraybuffer";

    var loader = this;

    request.onload = function() {
        loader.context.decodeAudioData(
            request.response,
            function(buffer) {
                if (!buffer) {
                    alert('error decoding file data: ' + url);
                    return;
                }
                loader.bufferList[index] = buffer;
                if (++loader.loadCount == loader.urlList.length)
                    loader.onload(loader.bufferList);
            }    
        );
    }

    request.onerror = function() {
        alert('BufferLoader: XHR error');        
    }

    request.send();
}

BufferLoader.prototype.load = function() {
    for (var i = 0; i < this.urlList.length; ++i)
        this.loadBuffer(this.urlList[i], i);
}

var context;
var bufferLoader;

function loadAndPlay() {
    try {
        context = new AudioContext();
    }
    catch(e) {
        alert("Web Audio API is not supported in this browser");
    }

    bufferLoader = new BufferLoader(
        context, 
        [ "sounds/kick.wav", "sounds/snare.wav", "sounds/hihat.wav", ],
        finishedLoading
    );

    bufferLoader.load();
}

function loadAndPlayStaggered() {
    try {
        context = new AudioContext();
    }
    catch(e) {
        alert("Web Audio API is not supported in this browser");
    }

    bufferLoader = new BufferLoader(
        context, 
        [ "sounds/kick.wav", "sounds/snare.wav", "sounds/hihat.wav", ],
        finishedLoadingStaggered
    );

    bufferLoader.load();
}

function finishedLoading(bufferList) {
    // Create three sources and buffers
    var kick = context.createBufferSource();
    var snare = context.createBufferSource();
    var hihat = context.createBufferSource();

    kick.buffer = bufferList[0];
    snare.buffer = bufferList[1];
    hihat.buffer = bufferList[2];
    
    kick.connect(context.destination);
    snare.connect(context.destination);
    hihat.connect(context.destination);	

	// Play them together
    kick.start(0);
    snare.start(0);
    hihat.start(0);	
}

function finishedLoadingStaggered(bufferList) {
    // Create three sources and buffers
    var kick = context.createBufferSource();
    var snare = context.createBufferSource();
    var hihat = context.createBufferSource();

    kick.buffer = bufferList[0];
    snare.buffer = bufferList[1];
    hihat.buffer = bufferList[2];
    
    kick.connect(context.destination);
    snare.connect(context.destination);
    hihat.connect(context.destination);	

	// Play them staggered
    kick.start(0);
    snare.start(0.125);
    hihat.start(0.25);	
}

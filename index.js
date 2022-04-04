document.getElementById('progress').indeterminate = true;

var worker = new Worker('worker.js')

worker.onmessage = function (message) {
    document.getElementById("progress").style.visibility = "hidden";
    document.getElementById("console").innerText = message.data;
};

const fileSelector = document.getElementById('file-selector');
fileSelector.addEventListener('change', (event) => {
    // event.target.files[0]
    if (event.target.files.length < 1) {
        return
    }
    document.getElementById("progress").style.visibility = "visible";
    document.getElementById("console").innerText = "Please wait while executing...";
    worker.postMessage(event.target.files[0]);
});




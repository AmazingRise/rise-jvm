document.getElementById('progress').indeterminate = true;

var worker = new Worker('worker.js')

worker.onmessage = function (message) {
    document.getElementById("progress").style.visibility = "hidden";
    document.getElementById("text").innerText = message.data + "\nProcess exited normally.";
    document.getElementById("file-selector").value = "";
};

const exec = function(file) {
    document.getElementById("progress").style.visibility = "visible";
    document.getElementById("text").innerText = "Please wait while executing...";
    worker.postMessage(file);
}

const fileSelector = document.getElementById('file-selector');
fileSelector.addEventListener('change', (event) => {
    if (event.target.files.length < 1) {
        return
    }
    exec(event.target.files[0]);
});

const load_demo = function () {
    fetch("Main.class").then(async (res)=>{
        worker.postMessage(await res.blob());
    })
}

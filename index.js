document.getElementById('progress').indeterminate = true;

var worker = new Worker('worker.js');

worker.onmessage = function (message) {
    if (typeof message.data === "boolean") {
        if (message.data) {
            document.getElementById("text").innerText = "Please wait while executing...";
        } else {
            document.getElementById("text").innerText = "Launching VM...";
        }
        return
    }
    document.getElementById("progress").style.visibility = "hidden";
    document.getElementById("text").innerText = message.data;
    document.getElementById("file-selector").value = "";
};

const exec = function(file) {
    console.clear();
    document.getElementById("progress").style.visibility = "visible";
    document.getElementById("text").innerText = "Reading the file...";
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
    document.getElementById("progress").style.visibility = "visible";
    document.getElementById("text").innerText = "Downloading byte code...";
    let name = document.getElementById("demo-select").value
    fetch("demo/" + name + ".class").then(async (res)=>{
        worker.postMessage(await res.blob());
    })

    fetch("demo/" + name + ".java").then(async (res)=>{
        document.getElementById("src").innerText = await res.text();
    })

    document.getElementById("demo-switch").style.visibility = "visible";
}

const toggle_src = function () {
    let a = "text";
    let b = "src";
    if (document.getElementById(b).style.display === "block") {
        [a, b] = [b, a];
    }
    document.getElementById(a).style.display = "none";
    document.getElementById(b).style.display = "block";
}

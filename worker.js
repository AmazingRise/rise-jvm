importScripts("wasm_exec.js");

var startTime;

onmessage = function(e) {

    const reader = new FileReader();
    reader.addEventListener('load', () => {
        postMessage(false);
        array = new Uint8Array(reader.result);
        const go = new self.Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
            postMessage(true);
            startTime = performance.now();
            try {
                const out = run(array, true);
                let timeDiff = Math.round(performance.now() - startTime);
                postMessage( out+ "\nProcess exited. Elapsed time: "+ timeDiff.toString() +"ms.");
            } catch {
                postMessage("Error happens. See console for details.\nProcess terminated.");
            }
        });
    });
    reader.readAsArrayBuffer(e.data);
}
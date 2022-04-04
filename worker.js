importScripts("wasm_exec.js");

onmessage = function(e) {
    const reader = new FileReader();
    reader.addEventListener('load', () => {
        array = new Uint8Array(reader.result);
        const go = new self.Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
            try {
                postMessage(run(array));
            } catch {
                postMessage("Error happens. Please open the console for details.");
            }

        });
    });
    reader.readAsArrayBuffer(e.data);
}
let channelURL = document.getElementById("channel-url");
let inputsToggled = false;

function startSpinner(id) {
    let spinner = document.getElementById(id);
    spinner.classList.remove("d-none");
}

function stopSpinner(id) {
    let spinner = document.getElementById(id);
    spinner.classList.add("d-none");
}

channelURL.addEventListener("change", (obj) => {
    if (channelURL.value === "") {
        toggleInputs();
        return
    }

    startSpinner("channel-metadata-spinner");

    let channel = {
        url: channelURL.value
    }

    const options = {
        method: "POST",
        body: JSON.stringify(channel),
        headers: new Headers({
            "Content-Type": "application/json"
        })
    };

    fetch("/api/channel/metadata", options)
        .then(res => res.json())
        .then(res => {
            // display all metadata and remove loading spinner
            let channelName = document.getElementById("channel-name")
            channelName.innerText = res.channelName

            let latestVideo = document.getElementById("channel-latest-video");
            latestVideo.innerText = res.latestVideo;
            latestVideo.href = res.latestVideoURL

            let latestVideoURL = document.getElementById("channel-latest-video-url")
            latestVideoURL.innerText = res.latestVideoURL
            latestVideoURL.href = res.latestVideoURL

            stopSpinner("channel-metadata-spinner")

            document.getElementById("channel-metadata").classList.remove("d-none")

            toggleInputs()
        });
})

function toggleInputs() {
    inputsToggled = !inputsToggled;
    document.getElementById("download-mode").disabled = !inputsToggled;
    document.getElementById("file-extension").disabled = !inputsToggled;
    document.getElementById("download-quality").disabled = !inputsToggled;
    document.getElementById("download-entire").disabled = !inputsToggled;
    document.getElementById("custom-ytdl-output").disabled = !inputsToggled;
    document.getElementById("custom-download-output").disabled = !inputsToggled;
    document.getElementById("custom-ytdl-command").disabled = !inputsToggled;
    document.getElementById("custom-download-output").disabled = !inputsToggled;
}

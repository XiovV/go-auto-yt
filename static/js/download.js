let url, downloadMode, fileExtension, resolution, downloadEntire, ytdlCommand;

function getDownloadPreferences() {
    url = document.getElementById("channel-url").value;

    let downloadModeElement = document.getElementById("download-mode");
    downloadMode = downloadModeElement.options[downloadModeElement.selectedIndex].value;

    let fileExtensionElement = document.getElementById("file-extension");
    fileExtension = fileExtensionElement.options[fileExtensionElement.selectedIndex].value;

    let resolutionElement = document.getElementById("download-quality");
    resolution = resolutionElement.options[resolutionElement.selectedIndex].value;

    downloadEntire = document.getElementById("download-entire").checked;

    let targetDownloadData = {
        url,
        downloadMode,
        fileExtension,
        resolution,
        downloadEntire,
        ytdlCommand,
    };
    console.log(targetDownloadData);
    return targetDownloadData;
}

function downloadTarget() {
    let downloadPreferences = getDownloadPreferences();
    const options = {
        method: "POST",
        body: JSON.stringify(downloadPreferences),
        headers: new Headers({
            "Content-Type": "application/json"
        })
    };
    console.log("REQ")
    fetch("/api/channel/add", options)
        .then(res => res.json())
        .then(res => {
            console.log(res)
        });

}
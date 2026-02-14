
// Handling file upload in frontend
async function upload() {
    // Get the first file from the file input
    const file = document.getElementById('fileInput').files[0];

    // Display message if file is not selected
    if(!file)
        return alert('Choose file')

    // Create FormData using the file
    const formData = new FormData();
    formData.append('file', file);

    // Sends POST /upload to backend with formData as body
    await fetch ('/upload', {
        method: 'POST',
        body: formData
    });

    // Fetches and displays updated file list
    loadFiles();
}


async function loadFiles() {
    const res = await fetch('/files');
    const files = await res.json();

    const list = document.getElementById('fileList')
    list.innerHTML = '';

    files.forEach(f => {
        const li = document.createElement('li');
        li.innerHTML = `${f} <a href="/download/${f}">Download</a>`;
        list.appendChild(li);
    });
}

loadFiles();
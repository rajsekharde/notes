async function upload() {
    const file = document.getElementById('fileInput').files[0];
    if(!file)
        return alert('Choose file')

    const formData = new FormData();
    formData.append('file', file);

    await fetch ('/upload', {
        method: 'POST',
        body: formData
    });

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
// fetchで画像一覧を取得し、画像を表示する
document.addEventListener('DOMContentLoaded', function() {
    fetch('/images')
        .then(response => response.json())
        .then(images => {
            const imagesList = document.getElementById('imagesList');
            images.forEach(image => {
                console.log(`/storage/${image.unique_string}`);
                const img = document.createElement('img');
                img.src = `/storage/${image.unique_string}`;
                img.alt = 'Uploaded Image';
                imagesList.appendChild(img);
            });
        });
});

// 画像をアップロードする
document.getElementById('uploadForm').addEventListener('submit', function(e) {
    e.preventDefault();

    const formData = new FormData(this);
    fetch('/upload', {
        method: 'POST',
        body: formData,
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        }
        throw new Error('Upload failed');
    })
    .then(data => {
        console.log('Upload successful', data);
        // 一覧を更新
    })
    .catch(error => {
        console.error('Error:', error);
    });
});

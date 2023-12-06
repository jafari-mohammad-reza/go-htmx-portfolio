function loadHTML(url, id) {
    fetch(url).then(function(response) {
        return response.text();
    }).then(function(html) {
        document.getElementById(id).innerHTML = html;
    }).catch(function(err) {
        console.warn('Something went wrong.', err);
    });
}
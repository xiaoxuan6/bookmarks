$(function () {
    NProgress.start();
    axios.get('/api/bookmarks').then(function (response) {
        let data = response.data;
        if (data.code !== 200) {
            NProgress.done();
            Notiflix.Notify.failure(data.msg);
        } else {
            let html = '';
            let items = data.data.Item;
            html += '<h1 class="text-2xl md:text-4xl font-bold mb-4">Welcome to BookMarks</h1>';
            html += '<div class="bg-white p-6 rounded-lg mb-4 shadow">';
            html += '<ul class="mb-4 flex">';
            $.each(items, function (index, item) {
                html += '<li class="mb-4 overflow-wrap"><a href="' + item['url'] + '" target="_blank">' + item['name'] + '</a></li>';
            });
            html += '</ul>';
            html += '</div>';
            document.getElementById('app').innerHTML = html;
            NProgress.done();
        }
    });
});
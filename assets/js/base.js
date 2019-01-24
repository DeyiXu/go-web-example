var setMenuActive = function () {
    var pathname = window.location.pathname;
    var $leftMenu = $("#leftMenu");
    var active = "active";
    var $liList = $leftMenu.find('li');
    $liList.removeClass(active);
    $liList.removeClass('menu-open');
    $liList.children('ul.treeview-menu').hide();

    var $checkedMenu = $leftMenu.find("a[href='" + pathname + "']");
    $checkedMenu.parents("#leftMenu li").addClass(active);
    $checkedMenu.parents("#leftMenu ul.treeview-menu").show();
};

var pjaxLoad = function () {
    $.pjax.defaults.timeout = 5000;
    $(document).pjax('a:not(.not-pjax)', '#pjax-container');
    $(document).on('submit', 'form[data-pjax]', function (event) {
        $.pjax.submit(event, '#pjax-container');
    });
    $(document).on('pjax:send', function () {
        console.debug("服务器请求发送~~~");
    });
    $(document).on('pjax:complete', function () {
        setMenuActive();
        if (window.PageComplete) {
            window.PageComplete();
            window.PageComplete = null;
        }
    });
    $(document).on('pjax:timeout', function (event) {
        console.debug("服务器请求超时~~~");
        event.preventDefault();
    });

    $(document).on('pjax:error', function (xhr, textStatus, error, options) {
        console.debug("pjaxerror:" + textStatus.status);
        if (textStatus.status == 500) {

        }
    });
    $(document).on('pjax:end', function () {
        console.debug("服务器请求结束~~~");
    });
};


$(document).ready(function () {
    pjaxLoad();
    setMenuActive();
    if (window.PageComplete) {
        window.PageComplete();
    }
});

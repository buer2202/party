// 滚动翻页插件，依赖jQuery WeUI扩展组件
(function ($) {
    $.fn.extend({
        "buerPages": function (options) {
            var that = this;
            var opts = $.extend({
                pageUrl: '/'
                , domHtml: function () {return '';}
            }, options); //使用jQuery.extend 覆盖插件默认参数

            show_page();

            var loading = false;  //状态标记
            $(document.body).infinite().on("infinite", function() {
                if (loading) return false;
                loading = true;
                show_page();
            });

            function show_page() {
                if (!opts.pageUrl) {
                    return null;
                }

                $.get(opts.pageUrl, function (data) {
                    var html;
                    for (var i = 0; i < data.List.length; i++) {
                        html = opts.domHtml(data.List[i]);
                        that.append(html);
                    }

                    loading = false;

                    // 如果是末页，就结束
                    if (data.LastPage) {
                        opts.pageUrl = null;
                    } else {
                        // 构造下一页的url
                        var url = opts.pageUrl.split('?');
                        opts.pageUrl = url[0] + '?p=' + (data.PageNo + 1);
                    }

                    if (!opts.pageUrl) {
                        $(document.body).destroyInfinite();
                        $('.weui-loadmore').text('已经到底了');
                    }
                });
            }

            return this;
        }
    });
})(jQuery);

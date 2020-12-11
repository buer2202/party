<div class="title">活动管理</div>
<div class="weui-btn-area">
    <a href="{{ urlfor `admin.PartyController.Create` }}" class="weui-btn weui-btn_plain-primary" id="submit">新建活动</a>
</div>

<div id="data-list"></div>
<div class="weui-loadmore">
    <i class="weui-loading"></i>
    <span class="weui-loadmore__tips">正在加载</span>
</div>

<script src="/static/js/buer_page.js"></script>
<script>
    // 滚动加载
    $('#data-list').buerPages({
        pageUrl: "{{ urlfor `admin.PartyController.DataList` }}",

        domHtml: function (data) {
            var html =  '<div class="weui-form-preview">'
                     +      '<div class="weui-form-preview__bd">'
                     +          '<div class="weui-form-preview__item">'
                     +              '<label class="weui-form-preview__label">活动题材</label>'
                     +              '<span class="weui-form-preview__value">' + data.Name + '</span>'
                     +          '</div>'
                     +          '<div class="weui-form-preview__item">'
                     +              '<label class="weui-form-preview__label">活动说明</label>'
                     +              '<span class="weui-form-preview__value">' + data.PartyDesc + '</span>'
                     +          '</div>'
                     +          '<div class="weui-form-preview__item">'
                     +              '<label class="weui-form-preview__label">活动确认</label>'
                     +              '<span class="weui-form-preview__value confirm_desc">' + data.ConfirmDesc + '</span>'
                     +          '</div>'
                     +      '</div>'
                     +      '<div class="weui-form-preview__ft">';

            if (!data.ConfirmDesc) {
                html +=         '<button class="weui-form-preview__btn weui-form-preview__btn_default party-confirm" data-id="' + data.Id + '">最终确认</button>';
            }

                html +=         '<button type="submit" class="weui-form-preview__btn weui-form-preview__btn_primary party-share" data-url_code="' + data.UrlCode + '">分享链接</button>'
                     +      '</div>'
                     +  '</div>';

            return html;
        }
    });

    // 活动确认
    $('#data-list').on('click', '.party-confirm', function () {
        var $this = $(this);

        layer.prompt({title: '输入活动最终确认方案！', formType: 2}, function (value, index) {
            layer.close(index);

            buer_post('{{ urlfor "admin.PartyController.Confirm" }}', {
                id: $this.data('id'),
                confirm_desc: value
            }, function (data, load) {
                layer.close(load)
                if (data.Status) {
                    layer.msg(data.Message, {icon: 6});
                    $this.parent().siblings('.weui-form-preview__bd').find('.confirm_desc').text(value);
                    $this.remove();
                } else {
                    layer.alert(data.Message, {icon: 5});
                }
            });
        });
    });
</script>

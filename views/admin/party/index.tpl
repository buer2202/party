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
                     +              '<label class="weui-form-preview__label">商品</label>'
                     +              '<span class="weui-form-preview__value">电动打蛋机</span>'
                     +          '</div>'
                     +          '<div class="weui-form-preview__item">'
                     +              '<label class="weui-form-preview__label">标题标题</label>'
                     +              '<span class="weui-form-preview__value">名字名字名字</span>'
                     +          '</div>'
                     +          '<div class="weui-form-preview__item">'
                     +              '<label class="weui-form-preview__label">标题标题</label>'
                     +              '<span class="weui-form-preview__value">很长很长的名字很长很长的名字很长很长的名字很长很长的名字很长很长的名字</span>'
                     +          '</div>'
                     +      '</div>'
                     +      '<div class="weui-form-preview__ft">'
                     +          '<a class="weui-form-preview__btn weui-form-preview__btn_default" href="javascript:">最终确认</a>'
                     +          '<button type="submit" class="weui-form-preview__btn weui-form-preview__btn_primary"'
                     +              'href="javascript:">分享链接</button>'
                     +      '</div>'
                     +  '</div>'

            return html;
        }
    });
</script>

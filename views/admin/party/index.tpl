<div class="title">
    <img src="/static/img/party.png" style="position:relative;top:3px;">
    活动管理
</div>
<div class="weui-btn-area">
    <a href="{{ urlfor `admin.PartyController.Create` }}" class="weui-btn weui-btn_plain-primary" id="submit">
        <img style="position:relative;top:8px;right:5px;" src="/static/img/new-party.png" />
        新建活动
    </a>
</div>

<div id="data-list"></div>
<div class="weui-loadmore">
    <i class="weui-loading"></i>
    <span class="weui-loadmore__tips">正在加载</span>
</div>

<div id="share-box" style="padding:20px;display:none;">
    <div>
        <img src="/static/img/share.png" style="position:relative;top:10px;">
        分享链接：
    </div>
    <input type="text" id="share-url" style="width:100%;border:0;color:#01AAED;outline:none;" />
    <div style="color:#5FB878;">（已复制，可直接去微信群粘贴啦！）</div>
</div>

<div id="confirm-window" style="display:none;">
    <div class="weui-cells__title">举办地点</div>
    <div class="weui-cells">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <input id="confirm_location" class="weui-input" type="text" placeholder="请输入地点名，如：艳阳天酒店(后湖店)" autocomplete="off">
            </div>
        </div>
    </div>

    <div class="weui-cells__title">具体描述</div>
    <div class="weui-cells weui-cells_form">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <textarea id="confirm_desc" class="weui-textarea" placeholder="请填写举办时间、详细地址、推荐交通工具等信息。" rows="8"></textarea>
            </div>
        </div>
    </div>
    <div class="weui-btn-area">
        <button class="weui-btn weui-btn_plain-primary" id="submit-confirm">
            <img style="position:relative;top:8px;right:5px;" src="/static/img/submit.png" />
            提交
        </button>
    </div>
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
                     +              '<label class="weui-form-preview__label">举办地点</label>'
                     +              '<span class="weui-form-preview__value confirm_location">' + data.ConfirmLocation + '</span>'
                     +          '</div>'
                     +          '<div class="weui-form-preview__item">'
                     +              '<label class="weui-form-preview__label">具体描述</label>'
                     +              '<span class="weui-form-preview__value confirm_desc">' + data.ConfirmDesc + '</span>'
                     +          '</div>'
                     +      '</div>'
                     +      '<div class="weui-form-preview__ft">'
                     +          '<button class="weui-form-preview__btn weui-form-preview__btn_primary party-redirect" data-url_code="' + data.UrlCode + '">查看</button>';
            if (!data.ConfirmDesc) {
                html +=         '<button class="weui-form-preview__btn weui-form-preview__btn_default party-confirm" data-id="' + data.Id + '">确认</button>';
            }

                html +=         '<button class="weui-form-preview__btn weui-form-preview__btn_primary party-share" data-url_code="' + data.UrlCode + '">分享</button>'
                     +      '</div>'
                     +  '</div>';

            return html;
        }
    });

    // 活动确认
    var $confirmButton;
    $('#data-list').on('click', '.party-confirm', function () {
        $confirmButton = $(this);

        layer.open({
            type: 1,
            shade: 0.3,
            title: "活动确认",
            area: ['90%', '480px'],
            content: $('#confirm-window')
        });
    });

    // 提交确认
    $('#submit-confirm').click(function () {
        buer_post('{{ urlfor "admin.PartyController.Confirm" }}', {
            id: $confirmButton.data('id'),
            confirm_location: $('#confirm_location').val(),
            confirm_desc: $('#confirm_desc').val()
        }, function (data, load) {
            layer.closeAll();
            if (data.Status) {
                layer.msg(data.Message, {icon: 6});
                $confirmButton.parent().siblings('.weui-form-preview__bd').find('.confirm_location').text($('#confirm_location').val());
                $confirmButton.parent().siblings('.weui-form-preview__bd').find('.confirm_desc').text($('#confirm_desc').val());
                $confirmButton.remove();
            } else {
                layer.alert("登录失效！", {icon: 5}, function () {
                    window.location.href = data.Content;
                });
            }
        });
    });

    // 查看活动
    $('#data-list').on('click', '.party-redirect', function () {
        $.get('{{ urlfor "admin.PartyController.ShareUrl" }}' + '?url_code=' + $(this).data('url_code'), function (data) {
            if (data.Status) {
                window.location.href = data.Content;
            } else {
                layer.alert("登录失效！", {icon: 5}, function () {
                    window.location.href = data.Content;
                });
            }
        }, 'json');
    });

    // 分享链接
    $('#data-list').on('click', '.party-share', function () {
        $.get('{{ urlfor "admin.PartyController.ShareUrl" }}' + '?url_code=' + $(this).data('url_code'), function (data) {
            if (data.Status) {
                layer.open({
                    type: 1,
                    shade: 0.3,
                    title: false,
                    area: ["90%", "150px"],
                    shadeClose: true,
                    content: $('#share-box')
                });

                // 赋值
                var $urlBox = $('#share-url');
                $urlBox.val(data.Content);

                var pwd = $urlBox[0];
                pwd.select(); // 选择对象
                document.execCommand("Copy"); // 执行浏览器复制命令
                pwd.blur(); // 取消焦点

            } else {
                layer.alert("登录失效！", {icon: 5}, function () {
                    window.location.href = data.Content;
                });
            }
        }, 'json');
    });
</script>

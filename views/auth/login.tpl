<div class="logo"><img src="/static/img/logo.jpg"></div>

<div class="weui-cells weui-cells_form">
    <div class="weui-cell">
        <div class="weui-cell__hd"><label class="weui-label">账号</label></div>
        <div class="weui-cell__bd">
            <input class="weui-input" type="text" id="account" placeholder="请输入登录账号" autocomplete="off">
        </div>
    </div>
    <div class="weui-cell">
        <div class="weui-cell__hd"><label class="weui-label">密码</label></div>
        <div class="weui-cell__bd">
            <input class="weui-input" type="password" id="password" placeholder="请输入登录密码" autocomplete="off">
        </div>
    </div>
</div>

<div class="weui-btn-area">
    <button class="weui-btn weui-btn_primary" type="button" id="submit">登录</button>
</div>
<script>
    $('#submit').click(function () {
        buer_post('{{ urlfor "AuthController.Login" }}', {
            account: $('#account').val(),
            password: $('#password').val()
        }, function (data, load) {
            if (data.Status) {
                window.location.href = '{{ urlfor "UserController.Get" }}';
            } else {
                layer.close(load);
                layer.msg(data.Message, {
                    icon: 5
                });
            }
        });
    });
</script>

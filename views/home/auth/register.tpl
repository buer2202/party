<h1 class="title">
    <img style="position:relative;top:8px;right:5px;" src="/static/img/register.png" />
    新用户注册
</h1>

<div class="weui-cells weui-cells_form">
    <div class="weui-cell">
        <div class="weui-cell__hd"><label class="weui-label">新账号</label></div>
        <div class="weui-cell__bd">
            <input class="weui-input" type="text" id="account" placeholder="请输入注册账号" autocomplete="off">
        </div>
    </div>
    <div class="weui-cell">
        <div class="weui-cell__hd"><label class="weui-label">密码</label></div>
        <div class="weui-cell__bd">
            <input class="weui-input" type="text" id="password" placeholder="请输入登录密码" autocomplete="off">
        </div>
    </div>
    <div class="weui-cell">
        <div class="weui-cell__hd"><label class="weui-label">昵称</label></div>
        <div class="weui-cell__bd">
            <input class="weui-input" type="text" id="nickname" placeholder="请输入您的昵称" autocomplete="off">
        </div>
    </div>
</div>

<div class="weui-btn-area">
    <button class="weui-btn weui-btn_plain-primary" type="button" id="submit">
        <img style="position:relative;top:8px;right:5px;" src="/static/img/submit.png" />
        立即注册
    </button>
    <a class="weui-btn weui-btn_plain-default" href="{{ urlfor `home.AuthController.LoginForm` }}">
        <img style="position:relative;top:8px;right:5px;" src="/static/img/login-form.png" />
        已有帐号，去登录！
    </a>
</div>
<script>
    $('#submit').click(function () {
        buer_post('{{ urlfor "home.AuthController.Regist" }}', {
            account: $('#account').val(),
            password: $('#password').val(),
            nickname: $('#nickname').val()
        }, function (data, load) {
            if (data.Status) {
                window.location.href = '{{ urlfor "admin.UserController.Get" }}';
            } else {
                layer.close(load);
                layer.msg(data.Message, {
                    icon: 5
                });
            }
        });
    });
</script>

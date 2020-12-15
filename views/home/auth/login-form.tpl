<h1 class="title">
    <img style="position:relative;top:8px;right:5px;" src="/static/img/login-form.png" />
    用户登录
</h1>

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
    <button class="weui-btn weui-btn_plain-primary" type="button" id="submit">
        <img style="position:relative;top:8px;right:5px;" src="/static/img/submit.png" />
        立即登录
    </button>
    <a class="weui-btn weui-btn_plain-default" href="{{ urlfor `home.AuthController.Register` }}">
        <img style="position:relative;top:8px;right:5px;" src="/static/img/register.png" />
        没有账号，去注册！
    </a>
</div>
<script>
    $('#submit').click(function () {
        buer_post('{{ urlfor "home.AuthController.Login" }}', {
            account: $('#account').val(),
            password: $('#password').val()
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

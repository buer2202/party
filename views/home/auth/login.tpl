<h1 class="title">用户登录</h1>

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

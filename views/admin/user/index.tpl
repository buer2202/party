<div class="title">
    <img style="position:relative;top:8px;right:5px;" src="/static/img/home.png" />
    {{ .authUser.Nickname }}，你好！
</div>
<div class="weui-grids">
    <a href="{{ urlfor `admin.PartyController.Index` }}" class="weui-grid js_grid">
        <div class="weui-grid__icon">
            <img src="/static/img/party.png">
        </div>
        <p class="weui-grid__label">
            活动管理
        </p>
    </a>
    <a href="{{ urlfor `admin.MemberController.Index` }}" class="weui-grid js_grid">
        <div class="weui-grid__icon">
            <img src="/static/img/members.png">
        </div>
        <p class="weui-grid__label">
            人员管理
        </p>
    </a>
    <a href="javascript:;" class="weui-grid js_grid" onclick="event.preventDefault();document.getElementById('logout-form').submit();">
        <div class="weui-grid__icon">
            <img src="/static/img/log-out.png">
        </div>
        <p class="weui-grid__label">
            退出登录
        </p>
        <form id="logout-form" action='{{ urlfor "home.AuthController.Logout" }}' method="POST" style="display: none;"></form>
    </a>
</div>

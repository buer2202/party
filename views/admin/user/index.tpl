<div class="title">Hello {{ .authUser.Nickname }}！</div>
<div class="weui-grids">
    <a href='{{ urlfor "PartyController.Create" }}' class="weui-grid js_grid">
        <div class="weui-grid__icon">
            <img src="/static/img/icon_nav_dialog.png" alt="">
        </div>
        <p class="weui-grid__label">
            发起
        </p>
    </a>
    <a href="" class="weui-grid js_grid">
        <div class="weui-grid__icon">
            <img src="/static/img/icon_nav_cell.png" alt="">
        </div>
        <p class="weui-grid__label">
            历史
        </p>
    </a>
    <a href="" class="weui-grid js_grid">
        <div class="weui-grid__icon">
            <img src="/static/img/icon_nav_article.png" alt="">
        </div>
        <p class="weui-grid__label">
            人员
        </p>
    </a>
    <a href="javascript:;" class="weui-grid js_grid" onclick="event.preventDefault();document.getElementById('logout-form').submit();">
        <div class="weui-grid__icon">
            <img src="/static/img/icon_nav_button.png" alt="">
        </div>
        <p class="weui-grid__label">
            退出
        </p>
        <form id="logout-form" action='{{ urlfor "AuthController.Logout" }}' method="POST" style="display: none;"></form>
    </a>
</div>

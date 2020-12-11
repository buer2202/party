<a href="{{ urlfor `admin.UserController.Get` }}" class="href-back">&lt;<i class="weui-icon-circle"></i></a>
<div class="title">成员管理</div>
<div class="weui-btn-area">
    <a href="{{ urlfor `admin.PartyController.Create` }}" class="weui-btn weui-btn_plain-primary" id="submit">新建成员</a>
</div>

<div class="weui-cells">
    {{range .dataList}}
    <div class="weui-cell weui-cell_swiped">
        <div class="weui-cell__bd">
            <div class="weui-cell">
                <div class="weui-cell__bd">{{ .Nickname }}</div>
                <div class="weui-cell__ft">左滑管理</div>
            </div>
        </div>
        <div class="weui-cell__ft">
            <span class="weui-swiped-btn weui-swiped-btn_warn delete-swipeout" href="javascript:">删除</span>
            <span class="weui-swiped-btn weui-swiped-btn_default close-swipeout" href="javascript:">编辑</span>
        </div>
    </div>
    {{end}}
</div>

<script>
    $('.weui-cell_swiped').swipeout()
</script>

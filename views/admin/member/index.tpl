<div class="title">
    <img src="/static/img/members.png" style="position:relative;top:5px;">
    成员管理
</div>
<div class="weui-btn-area">
    <button class="weui-btn weui-btn_plain-primary" id="add">
        <img style="position:relative;top:8px;right:5px;" src="/static/img/submit.png" />
        新建成员
    </button>
</div>

<div class="weui-cells">
    {{range .dataList}}
    <div class="weui-cell weui-cell_swiped">
        <div class="weui-cell__bd">
            <div class="weui-cell">
                <div class="weui-cell__bd">
                    <img src="/static/img/member.png" class="member-icon">
                    {{ .Nickname }}
                </div>
                <div class="weui-cell__ft">左滑管理</div>
            </div>
        </div>
        <div class="weui-cell__ft">
            <span class="weui-swiped-btn weui-swiped-btn_default edit"
                data-url="{{ urlfor `admin.MemberController.Update` `:id` .Id }}" data-nickname="{{ .Nickname }}">编辑</span>
            <span class="weui-swiped-btn weui-swiped-btn_warn destroy"
                data-url="{{ urlfor `admin.MemberController.Destroy` `:id` .Id }}">删除</span>
        </div>
    </div>
    {{end}}
</div>

<script>
    $('.weui-cell_swiped').swipeout()

    // 添加
    $('#add').click(function () {
        layer.prompt({
            formType: 0,
            title: '添加新成员'
        }, function (value, index, elem) {
            layer.close(index);
            buer_post('{{ urlfor "admin.MemberController.Store" }}', {nickname: value}, function (data, load) {
                layer.close(load)
                if (data.Status) {
                    layer.alert('操作成功', {icon: 6}, function () {
                        window.location.reload();
                    });
                } else {
                    if (data.Content) {
                        layer.alert("登录失效！", {icon: 5}, function () {
                            window.location.href = data.Content;
                        });
                    } else {
                        layer.alert(data.Message);
                    }
                }
            });
        });
    });

    // 编辑
    $('.edit').click(function () {
        var url = $(this).data('url');

        layer.prompt({
            formType: 0,
            value: $(this).data('nickname'),
            title: '修改成员'
        }, function (value, index, elem) {
            layer.close(index);
            buer_post(url, {nickname: value}, function (data, load) {
                layer.close(load)
                if (data.Status) {
                    layer.alert('操作成功', {icon: 6}, function () {
                        window.location.reload();
                    });
                } else {
                    if (data.Content) {
                        layer.alert("登录失效！", {icon: 5}, function () {
                            window.location.href = data.Content;
                        });
                    } else {
                        layer.alert(data.Message);
                    }
                }
            });
        });
    });

    // 删除
    $('.destroy').click(function () {
        var url = $(this).data('url');

        layer.confirm("确定删除吗？", function () {
            buer_post(url);
        });
    });
</script>

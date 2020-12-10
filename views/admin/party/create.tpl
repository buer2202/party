<div class="title">新建活动</div>

<div class="weui-cells__title">活动题材</div>
<div class="weui-cells">
    <div class="weui-cell">
        <div class="weui-cell__bd">
            <input id="name" class="weui-input" type="text" placeholder="请输入">
        </div>
    </div>
</div>

<div class="weui-cells__title">活动描述</div>
<div class="weui-cells weui-cells_form">
    <div class="weui-cell">
        <div class="weui-cell__bd">
            <textarea id="party_desc" class="weui-textarea" placeholder="请输入" rows="8"></textarea>
        </div>
    </div>
</div>
<div class="weui-btn-area">
    <button class="weui-btn weui-btn_primary" id="submit">提交</button>
</div>
<script>
    $('#submit').click(function () {
        buer_post('{{ urlfor "admin.PartyController.Store" }}', {
            account: $('#account').val(),
            password: $('#password').val()
        }, function (data, load) {
            if (data.Status) {
                window.location.reload();
            } else {
                layer.close(load);
                layer.msg(data.Message, {
                    icon: 5
                });
            }
        });
    });
</script>

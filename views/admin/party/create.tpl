<div class="title">
    <img src="/static/img/new-party.png" style="position:relative;top:6px;">
    新建活动
</div>

<div class="weui-cells__title">活动题材</div>
<div class="weui-cells">
    <div class="weui-cell">
        <div class="weui-cell__bd">
            <input id="party_name" class="weui-input" type="text" placeholder="请输入" autocomplete="off">
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
    <button class="weui-btn weui-btn_plain-primary" id="submit">
        <img style="position:relative;top:8px;right:5px;" src="/static/img/submit.png" />
        提交
    </button>
</div>
<script>
    $('#submit').click(function () {
        buer_post('{{ urlfor "admin.PartyController.Store" }}', {
            name: $('#party_name').val(),
            party_desc: $('#party_desc').val()
        }, function (data, load) {
            if (data.Status) {
                layer.alert("创建成功！", function () {
                    window.location.href = '{{ urlfor "admin.PartyController.Index" }}';
                });
            } else {
                layer.close(load);
                layer.msg(data.Message, {
                    icon: 5
                });
            }
        });
    });
</script>

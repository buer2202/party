<h1 class="title">{{ .party.Name }}</h1>

<div class="section">当前活动：</div>
<div class="weui-cells__tips">{{ .party.PartyDesc }}</div>

<div class="weui-cells weui-cells_form">
    <div class="weui-cell weui-cell_select weui-cell_select-after">
        <div class="weui-cell__hd">
            <label for="" class="weui-label">我是</label>
        </div>
        <div class="weui-cell__bd">
            <select class="weui-select" id="member_id">
                <option value="">请选择</option>
                {{ range .members }}
                <option value="{{ .Id }}">{{ .Nickname }}</option>
                {{ end }}
            </select>
        </div>
    </div>
    <div class="weui-cell weui-cell_select weui-cell_select-after">
        <div class="weui-cell__hd">
            <label for="" class="weui-label">人数</label>
        </div>
        <div class="weui-cell__bd">
            <select class="weui-select" id="join_people_num">
                <option value="1">1人</option>
                <option value="2">2人</option>
                <option value="3">3人</option>
                <option value="4">4人</option>
                <option value="5">5人</option>
            </select>
        </div>
    </div>
    <div class="weui-cell">
        <div class="weui-cell__hd"><label class="weui-label">时间</label></div>
        <div class="weui-cell__bd">
            <input id="can_join_date" class="weui-input" type="text" placeholder="点日历选有空时间">
        </div>
    </div>
</div>

<div id="inline-calendar"></div>

<div class="weui-btn-area">
    <button class="weui-btn weui-btn_primary" id="submit">提交</button>
</div>

<div class="section">报名情况：</div>
<table class="my-table">
    <tr>
        <td>日期\人员</td>
        {{ range .members}}
        <td>{{ .Nickname }}</td>
        {{ end }}
    </tr>
    {{ range $k, $date := .partyMemberDate }}
    <tr>
        <td>{{ $date.CanJoinDate }}</td>
        {{ range $.members}}
        <td class="join-show {{ .Id }}-{{ $date.CanJoinDate }}"></td>
        {{ end }}
    </tr>
    {{ end }}
</table>

<div class="section">确认方案：</div>
<div class="confirm-desc">{{ .party.ConfirmDesc }}</div>

<script>
    $("#inline-calendar").calendar({
        container: "#inline-calendar",
        input: "#can_join_date",
        multiple: true,
        value: []
    });

    // 提交
    $('#submit').click(function () {
        buer_post('{{ urlfor "PartyController.Post" ":urlCode" .party.UrlCode }}', {
            member_id: $('#member_id').val(),
            join_people_num: $('#join_people_num').val(),
            can_join_date: $('#can_join_date').val()
        });
    });

    $.get('{{ urlfor "PartyController.PartyMembers" ":urlCode" .party.UrlCode }}', function (data) {
        if (data) {
            var tdClassName;
            for (const i in data) {
                tdClassName = '.' + data[i].MemberId + '-' + data[i].CanJoinDate;
                $(tdClassName).addClass('can-join').text(data[i].JoinPeopleNum);
            }
        }
    }, 'json');
</script>

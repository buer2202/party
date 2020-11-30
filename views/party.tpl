<!DOCTYPE html>
<html>

<head>
    <title>Party2202</title>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=no">
    <link rel="stylesheet" href="/static/plug-in/jquery-weui/lib/weui.min.css">
    <link rel="stylesheet" href="/static/plug-in/jquery-weui/css/jquery-weui.min.css">
    <link rel="stylesheet" href="/static/css/mobile.css">
</head>

<body>
    <div class="title">{{ .party.Name }}</div>

    <div class="weui-cells__tips">
        当前活动：{{ .party.PartyDesc }}
    </div>

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

    <div class="result">
        报名情况：
    </div>
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

    <div class="footer">
        <img src="/static/img/logo.jpg" />
        <div class="weui-msg__desc">Power by 不耳 2020.11</div>
    </div>

    <script src="/static/plug-in/jquery-weui/lib/jquery-2.1.4.js"></script>
    <script src="/static/plug-in/jquery-weui/js/jquery-weui.js"></script>
    <script>
        $("#inline-calendar").calendar({
            container: "#inline-calendar",
            input: "#can_join_date",
            multiple: true,
            value: []
        });

        // 提交
        $('#submit').click(function () {
            $.post('/party/{{ .urlCode }}', {
                member_id: $('#member_id').val(),
                join_people_num: $('#join_people_num').val(),
                can_join_date: $('#can_join_date').val()
            }, function (data) {
                if (data) {
                    $.toast("操作成功", function () {
                        window.location.reload();
                    });
                } else {
                    $.toast("提交失败");
                }
            }, 'json');
        });

        $.get('/party/{{ .party.UrlCode }}/party-members', function (data) {
            if (data) {
                var tdClassName;
                for (const i in data) {
                    tdClassName = '.' + data[i].MemberId + '-' + data[i].CanJoinDate;
                    $(tdClassName).addClass('can-join').text(data[i].JoinPeopleNum);
                }
            }
        }, 'json');
    </script>
</body>

</html>

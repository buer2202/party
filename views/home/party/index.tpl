<h1 class="title">
    <img src="/static/img/share.png" style="position:relative;top:10px;">
    {{ .party.Name }}
</h1>

<div class="section" style="margin-top: 0;">
    <img src="/static/img/party-desc.png" style="position:relative;top:10px;">
    活动描述：
</div>
<div class="weui-cells__tips">{{ .party.PartyDesc }}</div>

{{ if .party.ConfirmLocation }}
    <div class="section">
        <img src="/static/img/confirm-desc.png" style="position:relative;top:10px;">
        确认方案：
    </div>

    <script type="text/javascript" src="https://api.map.baidu.com/api?v=1.0&type=webgl&ak=CC1448ce94eb75c90d9c2192d2a9a42f"></script>
    <div id="container" style="height:250px;width:100%;"></div>
    <script type="text/javascript">
        var map = new BMapGL.Map("container");

        //智能搜索
        var local = new BMapGL.LocalSearch(map, {
            onSearchComplete: function () {
                //获取第一个智能搜索的结果
                var pp = local.getResults().getPoi(0).point;
                // console.log('经度：'+pp.lng, '纬度：'+pp.lat);

                // 定位
                map.centerAndZoom(pp, 18);
                //添加标注
                var myIcon = new BMapGL.Icon("/static/favicon.ico", new BMapGL.Size(30, 30));
                var myMarker = new BMapGL.Marker(pp, {icon: myIcon});
                map.addOverlay(myMarker);

                // 添加说明窗口
                var opts = {
                    width: 200,
                    height: 100,
                    title: '{{ .party.ConfirmLocation }}'
                };
                var infoWindow = new BMapGL.InfoWindow('{{ .party.ConfirmDesc }}', opts);
                map.openInfoWindow(infoWindow, pp);
                myMarker.addEventListener('click', function () {
                    map.openInfoWindow(infoWindow, pp);
                });
            }
        });
        local.search("{{ .party.ConfirmLocation }}");
    </script>

{{ else }}
    <div class="section">
        <img src="/static/img/join-form.png" style="position:relative;top:10px;">
        开始报名：<span class="weui-cells__tips">（重复报名会覆盖前一次）</span>
    </div>
    <div class="weui-cells weui-cells_form">
        <div class="weui-cell weui-cell_select weui-cell_select-after">
            <div class="weui-cell__hd">
                <label for="" class="weui-label">我是</label>
            </div>
            <div class="weui-cell__bd">
                <select class="weui-select" id="member_id">
                    <option value="">请选择</option>
                    {{ range .allMembers }}
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
        <button class="weui-btn weui-btn_plain-primary" id="submit">
            <img src="/static/img/join.png" style="position:relative;top:10px;">
            立即报名
        </button>
    </div>
{{ end }}

<div class="section">
    <img src="/static/img/joined.png" style="position:relative;top:10px;">
    报名情况：
</div>
<table class="my-table">
    <tr>
        <td width="100"><img src="/static/img/joined-table.png"></td>
        {{ range .joinedMembers}}
        <td>{{ .MemberNickname }}</td>
        {{ end }}
    </tr>
    {{ range $k, $date := .partyMemberDate }}
    <tr>
        <td>{{ substr $date.CanJoinDate 5 5 }}</td>
        {{ range $.joinedMembers}}
        <td class="join-show {{ .MemberNickname }}-{{ $date.CanJoinDate }}"></td>
        {{ end }}
    </tr>
    {{ end }}
</table>

<script>
    $("#inline-calendar").calendar({
        container: "#inline-calendar",
        input: "#can_join_date",
        multiple: true,
        value: []
    });

    // 提交
    $('#submit').click(function () {
        buer_post('{{ urlfor "home.PartyController.Post" ":urlCode" .party.UrlCode }}', {
            member_id: $('#member_id').val(),
            join_people_num: $('#join_people_num').val(),
            can_join_date: $('#can_join_date').val()
        });
    });

    $.get('{{ urlfor "home.PartyController.PartyMembers" ":urlCode" .party.UrlCode }}', function (data) {
        if (data) {
            var tdClassName;
            for (const i in data) {
                tdClassName = '.' + data[i].MemberNickname + '-' + data[i].CanJoinDate;
                $(tdClassName).addClass('can-join').text(data[i].JoinPeopleNum);
            }
        }
    }, 'json');
</script>

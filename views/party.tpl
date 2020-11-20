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
    <header class="demos-header">
        <h1 class="demos-title">{{ .party.Name }}</h1>
    </header>

    <div class="weui-cells__tips">
        当前活动：{{ .party.PartyDesc }} {{ .members }}
    </div>
    <div id="inline-calendar"></div>

    <div class="weui-cells weui-cells_form">
        <div class="weui-cell">
            <div class="weui-cell__hd"><label class="weui-label">已选日期</label></div>
            <div class="weui-cell__bd">
                <input id="input-choosed" class="weui-input" type="text" placeholder="点击可参加日期">
            </div>
        </div>
    </div>

    <script src="/static/plug-in/jquery-weui/lib/jquery-2.1.4.js"></script>
    <script src="/static/plug-in/jquery-weui/js/jquery-weui.js"></script>
    <script>
        $("#inline-calendar").calendar({
            container: "#inline-calendar",
            input: "#input-choosed",
            multiple: true,
            value: []
        });
    </script>
</body>

</html>

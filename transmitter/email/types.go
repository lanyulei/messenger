package email

/*
  @Author : lanyulei
  @Desc :
*/

const templateData = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>ferry</title>
</head>
<body>
    <br>
    工单信息如下：
    <br>
    <br>
    <table>
        <tr>
            <td style="text-align: right">标题：</td>
            <td>{{ .Title }}</td>
        </tr>
        <tr>
            <td style="text-align: right">申请人：</td>
            <td>{{ .Creator }}</td>
        </tr>
        <tr>
            <td style="text-align: right">优先级：</td>
            <td>{{ .Priority }}</td>
        </tr>
        <tr>
            <td style="text-align: right">申请时间：</td>
            <td>{{ .CreatedAt }}</td>
        </tr>
    </table>
</body>
<style>
    table {
        border: 1px solid #ccc;
        border-collapse:collapse;
    }
    td {
        padding: 10px 15px 10px 15px;
        border: 1px solid #ccc;
    }
</style>
</html>`

$(document).ready(function () {
    $("#uploadfile").change(function () {
        $("#uploadfile").upload('/image/upload2', function (data) {
            var urls = data.split('|');
            urls.forEach(element => {
                if (!!element)
                    $("#picdiv").html("<img style='height:60px;margin-left:10px;' src='" + element + "'>");
            });
        });
    });
});
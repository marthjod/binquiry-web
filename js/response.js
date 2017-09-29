var clearFields = function () {
    "use strict";

    $(".input-group").removeClass("has-error");
    $("#tag-response").find("span").remove();
};

var prettifyJson = function (json) {
    "use strict";

    return JSON.stringify(json, null, 2);
};

var isError = function (json) {
    return json && typeof json === 'object' && json.hasOwnProperty("error");
};

var updateResponseBlock = function (response) {
    "use strict";

    var responseBlock = $("code#response"),
        text = "";

    clearFields();

    if (isError(response)) {
        $(".input-group").addClass("has-error");
        text = response.error;
    } else {
        text = prettifyJson(response);
    }

    $(responseBlock).text(text);
};

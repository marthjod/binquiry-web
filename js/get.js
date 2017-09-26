const defaultQuery = "orð";

var submitQuery = function() {
    "use strict";

    var query = $("input#query").val();
    query = query === "" ? defaultQuery : query;

    $.ajax({
        url: getQueryUrl(query),
        success: function (result) {
            updateResponseBlock(result);
        },
        fail: function (result) {
            updateResponseBlock(result);
        }
    })
};

var getQueryUrl = function(query) {
    "use strict";

    return "/?q=" + query;
};

var prettifyJson = function(json) {
    "use strict";

    return JSON.stringify(json, null, 2);
};

var isError = function (json) {
    return json && typeof json === 'object' && json.hasOwnProperty("error");
};

var updateResponseBlock = function(response) {
    "use strict";

    var block = $("code#response"),
        text = "";

    if (isError(response)) {
        text = response["error"];
    } else {
        text = prettifyJson(response);
    }

    $(block).text(text);
};

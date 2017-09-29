const defaultQuery = "orð";

var submitQuery = function () {
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

var getQueryUrl = function (query) {
    "use strict";

    return "/?q=" + query;
};

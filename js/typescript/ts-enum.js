/* TS Code
enum HttpCode {
    // 成功
    '200_OK' = 200,
    // 已生成了新的资源
    '201_Created' = 201,
    // 请求稍后会被处理
    '202_Accepted' = 202,
    // 资源已经不存在
    '204_NoContent' = 204,
    // 被请求的资源有一系列可供选择的回馈信息
    '300_MultipleChoices' = 300,
    // 永久性转移
    '301_MovedPermanently' = 301,
    // 暂时性转移
    '302_MoveTemporarily' = 302,
}

HttpCode['200_OK']
HttpCode[200]
*/

/* compiled into following js code */

"use strict";
var HttpCode;
(function (HttpCode) {
    /** 成功 */
    HttpCode[HttpCode["200_OK"] = 200] = "200_OK";
    /** 已生成了新的资源 */
    HttpCode[HttpCode["201_Created"] = 201] = "201_Created";
    /** 请求稍后会被处理 */
    HttpCode[HttpCode["202_Accepted"] = 202] = "202_Accepted";
    /** 资源已经不存在 */
    HttpCode[HttpCode["204_NoContent"] = 204] = "204_NoContent";
    /** 被请求的资源有一系列可供选择的回馈信息 */
    HttpCode[HttpCode["300_MultipleChoices"] = 300] = "300_MultipleChoices";
    /** 永久性转移 */
    HttpCode[HttpCode["301_MovedPermanently"] = 301] = "301_MovedPermanently";
    /** 暂时性转移 */
    HttpCode[HttpCode["302_MoveTemporarily"] = 302] = "302_MoveTemporarily";
})(HttpCode || (HttpCode = {}));

console.log(HttpCode['200_OK']);
console.log(HttpCode[200]);

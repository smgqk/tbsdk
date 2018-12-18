package tbsdk

// TaobaoFenxiaoProductQuantityUpdateRequest 修改产品库存信息，支持全量修改以及增量修改两种方式
type TaobaoFenxiaoProductQuantityUpdateRequest struct {
    
    /* product_id required产品ID */
    product_id int64 `json:"product_id";xml:"product_id"`
    
    /* properties optionalsku属性值，产品有sku时填写，多个sku用,分隔。为空时默认该产品无sku，则只修改产品的库存。 */
    properties string `json:"properties";xml:"properties"`
    
    /* quantity required库存修改值。产品有sku时，与sku属性顺序对应，用,分隔。产品无sku时，只写库存值。
当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0 */
    quantity string `json:"quantity";xml:"quantity"`
    
    /* _type optional库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0 */
    _type int64 `json:"type";xml:"type"`
    
}

func (req *TaobaoFenxiaoProductQuantityUpdateRequest) GetAPIName() string {
	return "taobao.fenxiao.product.quantity.update"
}

// TaobaoFenxiaoProductQuantityUpdateResponse 修改产品库存信息，支持全量修改以及增量修改两种方式
type TaobaoFenxiaoProductQuantityUpdateResponse struct {
    
    /* created Basic操作时间 */
    created Date `json:"created";xml:"created"`
    
    /* result Basic操作结果 */
    result bool `json:"result";xml:"result"`
    
}

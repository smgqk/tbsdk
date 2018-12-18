package tbsdk

// TmallChannelProductsGetRequest 查询供应商的产品数据。 

* 入参传入pids将优先查询，即只按这个条件查询。 
*入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条) 
* 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。 
* 入参fields传入images将查询多图数据，不传只返回主图数据。 
* 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据） 
* 查询结果按照产品发布时间倒序，即时间近的数据在前。
* 传入channel 渠道，会只返回相应渠道的产品
type TmallChannelProductsGetRequest struct {
    
    /* top_query_product_d_o optionaltop_query_product_d_o */
    top_query_product_d_o TopQueryProductDo `json:"top_query_product_d_o";xml:"top_query_product_d_o"`
    
}

func (req *TmallChannelProductsGetRequest) GetAPIName() string {
	return "tmall.channel.products.get"
}

// TmallChannelProductsGetResponse 查询供应商的产品数据。 

* 入参传入pids将优先查询，即只按这个条件查询。 
*入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条) 
* 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。 
* 入参fields传入images将查询多图数据，不传只返回主图数据。 
* 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据） 
* 查询结果按照产品发布时间倒序，即时间近的数据在前。
* 传入channel 渠道，会只返回相应渠道的产品
type TmallChannelProductsGetResponse struct {
    
    /* products Object Array产品对象记录集 */
    products TopProductDO `json:"products";xml:"products"`
    
    /* total_results Basic查询结果记录数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}
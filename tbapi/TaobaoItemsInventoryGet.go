package tbsdk

// TaobaoItemsInventoryGetRequest 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤 
只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取
type TaobaoItemsInventoryGetRequest struct {
    
    /* auction_type optional商品类型：a-拍卖,b-一口价 */
    auction_type string `json:"auction_type";xml:"auction_type"`
    
    /* banner optional分类字段。可选值:<br>
regular_shelved(定时上架)<br>
never_on_shelf(从未上架)<br>
off_shelf(我下架的)<br>
<font color='red'>for_shelved(等待所有上架)<br>
sold_out(全部卖完)<br>
violation_off_shelf(违规下架的)<br>
默认查询for_shelved(等待所有上架)这个状态的商品<br></font>
注：for_shelved(等待所有上架)=regular_shelved(定时上架)+never_on_shelf(从未上架)+off_shelf(我下架的) */
    banner string `json:"banner";xml:"banner"`
    
    /* cid optional商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到 */
    cid int64 `json:"cid";xml:"cid"`
    
    /* end_modified optional商品结束修改时间 */
    end_modified Date `json:"end_modified";xml:"end_modified"`
    
    /* fields required需返回的字段列表。可选值为 Item 商品结构体中的以下字段： 
approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru, list_time,price,has_discount,has_invoice,has_warranty,has_showcase, modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。<br> 
不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get。 */
    fields string `json:"fields";xml:"fields"`
    
    /* has_discount optional是否参与会员折扣。可选值：true，false。默认不过滤该条件 */
    has_discount bool `json:"has_discount";xml:"has_discount"`
    
    /* is_ex optional商品是否在外部网店显示 */
    is_ex bool `json:"is_ex";xml:"is_ex"`
    
    /* is_taobao optional商品是否在淘宝显示 */
    is_taobao bool `json:"is_taobao";xml:"is_taobao"`
    
    /* order_by optional排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc */
    order_by string `json:"order_by";xml:"order_by"`
    
    /* page_no optional页码。取值范围:大于零小于等于101的整数;默认值为1，即返回第一页数据。当页码超过101页时系统就会报错，故请大家在用此接口获取数据时尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品。 */
    page_no int64 `json:"page_no";xml:"page_no"`
    
    /* page_size optional每页条数。取值范围:大于零的整数;最大值：200；默认值：40。 */
    page_size int64 `json:"page_size";xml:"page_size"`
    
    /* q optional搜索字段。搜索商品的title。 */
    q string `json:"q";xml:"q"`
    
    /* seller_cids optional卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>) */
    seller_cids string `json:"seller_cids";xml:"seller_cids"`
    
    /* start_modified optional商品起始修改时间 */
    start_modified Date `json:"start_modified";xml:"start_modified"`
    
}

func (req *TaobaoItemsInventoryGetRequest) GetAPIName() string {
	return "taobao.items.inventory.get"
}

// TaobaoItemsInventoryGetResponse 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤 
只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取
type TaobaoItemsInventoryGetResponse struct {
    
    /* items Object Array搜索到底商品列表，具体字段根据设定的fields决定，不包括desc,stuff_status字段 */
    items Item `json:"items";xml:"items"`
    
    /* total_results Basic搜索到符合条件的结果总数 */
    total_results int64 `json:"total_results";xml:"total_results"`
    
}

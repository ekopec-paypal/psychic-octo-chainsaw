package scratch

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	ExchangeTag = "happyExchange"
)

type PaymentDetails struct {
	AVSResultCode     string `json:"avs_result_code"`
	CreditCardBin     string `json:"credit_card_bin"`
	CVVResultCode     string `json:"cvv_result_code"`
	CreditCardCompany string `json:"credit_card_company"`
	CreditCardNumber  string `json:"credit_card_number"`
}

type Transaction struct {
	ID                int64           `json:"id,omitempty"`
	OrderID           int64           `json:"order_id,omitempty"`
	CreatedAt         string          `json:"created_at,omitempty"`
	Authorization     string          `json:"authorization,omitempty"`
	Amount            string          `json:"amount,omitempty"`
	MaximumRefundable string          `json:"maximum_refundable,omitempty"`
	Currency          string          `json:"currency,omitempty"`
	Gateway           string          `json:"gateway,omitempty"`
	SourceName        string          `json:"source_name,omitempty"`
	PaymentDetails    *PaymentDetails `json:"payment_details,omitempty"`
	Kind              string          `json:"kind,omitempty"`
	Status            string          `json:"status,omitempty"`
	ErrorCode         string          `json:"error_code,omitempty"`
	Test              bool            `json:"test,omitempty"`
	ParentID          int64           `json:"parent_id,omitempty"`
	Message           string          `json:"message,omitempty"`
	Source            string          `json:"source,omitempty"`
	// This is all custom garbage, so we're ignoring for now
	//UserID         int64           `json:"user_id,omitempty"`
	//Authorization  string          `json:"authorization,omitempty"`
	//DeviceID       ForcedString    `json:"device_id,omitempty"`
	//Receipt        struct{}       `json:"receipt"`
	//LocationID
}

type RefundItem struct {
	ID                      int64             `json:"id,omitempty"`
	LineItem                *OrderItem        `json:"line_item,omitempty"`
	LineItemID              int64             `json:"line_item_id"`
	Quantity                int               `json:"quantity"`
	RestockType             string            `json:"restock_type,omitempty"`
	LocationID              int64             `json:"location_id,omitempty"`
	Price                   ForcedFloat       `json:"price,omitempty"`
	Subtotal                ForcedFloat       `json:"subtotal,omitempty"`
	SubtotalSet             *MultiCurrencySet `json:"subtotal_set,omitempty"`
	TotalTax                ForcedFloat       `json:"total_tax,omitempty"`
	TotalTaxSet             *MultiCurrencySet `json:"total_tax_set,omitempty"`
	DiscountedPrice         ForcedFloat       `json:"discounted_price,omitempty"`
	DiscountedTotalPrice    ForcedFloat       `json:"discounted_total_price,omitempty"`
	TotalCartDiscountAmount ForcedFloat       `json:"total_cart_discount_amount,omitempty"`
	RefundOptionID          string            `json:"-"`
}

type CurrencyAmount struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

type MultiCurrencySet struct {
	PresentmentMoney CurrencyAmount `json:"presentment_money"`
	ShopMoney        CurrencyAmount `json:"shop_money"`
}

type RefundShipping struct {
	FullRefund bool        `json:"full_refund"`
	Amount     ForcedFloat `json:"amount,omitempty"`
}

type Refund struct {
	ID                int64              `json:"id,omitempty"`
	Currency          string             `json:"currency"`
	CreatedAt         string             `json:"created_at,omitempty"`
	Notify            bool               `json:"notify,omitempty"`
	Note              string             `json:"note,omitempty"`
	DiscrepancyReason string             `json:"discrepancy_reason"`
	LineItems         []*RefundItem      `json:"refund_line_items"`
	Transactions      []*Transaction     `json:"transactions,omitempty"`
	UserID            int64              `json:"user_id,omitempty"`
	Shipping          *RefundShipping    `json:"shipping,omitempty"`
	OrderAdjustments  []*OrderAdjustment `json:"order_adjustments"`
}

type OrderAdjustment struct {
	ID           int64             `json:"id,omitempty"`
	OrderID      int64             `json:"order_id,omitempty"`
	RefundID     int64             `json:"refund_id,omitempty"`
	Amount       string            `json:"amount,omitempty"`
	TaxAmount    string            `json:"tax_amount,omitempty"`
	Kind         string            `json:"kind,omitempty"`
	Reason       string            `json:"reason,omitempty"`
	AmountSet    *MultiCurrencySet `json:"amount_set,omitempty"`
	TaxAmountSet *MultiCurrencySet `json:"tax_amount_set,omitempty"`
}

type NoteAttribute struct {
	Name  ForcedString `json:"name"`
	Value ForcedString `json:"value"`
}

type DiscountApplication struct {
	Type            string      `json:"type"`              //"manual",
	Title           string      `json:"title"`             //"custom discount",
	Description     string      `json:"description"`       //"customer deserved it",
	Value           ForcedFloat `json:"value"`             //"2.0",
	ValueType       string      `json:"value_type"`        //"fixed_amount",
	Allocation      string      `json:"allocation_method"` //"across",
	TargetSelection string      `json:"target_selection"`  //"explicit",
	TargetType      string      `json:"target_type"`       //"line_item"
}

type DiscountCode struct {
	Code   string `json:"code"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type InvoicePaidRequestBody struct {
	ID         int64  `json:"id,omitempty"`
	OrderID    int64  `json:"order_id,omitempty"`
	Status     string `json:"status,omitempty"`
	TotalPrice string `json:"total_price,omitempty"`
}

type Order struct {
	ID                   int64                  `json:"id,omitempty"`
	OrderNumber          int64                  `json:"order_number,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	CreatedAt            string                 `json:"created_at,omitempty"`
	Customer             *Customer              `json:"customer,omitempty"`
	Email                string                 `json:"email,omitempty"`
	LineItems            []*OrderItem           `json:"line_items,omitempty"`
	Transactions         []*Transaction         `json:"transactions,omitempty"`
	Fulfillments         []Fulfillment          `json:"fulfillments,omitempty"`
	FulfillmentStatus    string                 `json:"fulfillment_status,omitempty"`
	DiscountApplications []*DiscountApplication `json:"discount_applications"`
	DiscountCodes        []*DiscountCode        `json:"discount_codes"`
	Refunds              []Refund               `json:"refunds,omitempty"`
	Note                 string                 `json:"note,omitempty"`
	NoteAttributes       []NoteAttribute        `json:"note_attributes,omitempty"`
	Tags                 string                 `json:"tags,omitempty"`
	SourceName           string                 `json:"source_name,omitempty"`
	BillingAddress       *ShippingAddress       `json:"billing_address,omitempty"`
	ShippingAddress      *ShippingAddress       `json:"shipping_address,omitempty"`
	Currency             string                 `json:"currency,omitempty"`
	TotalDiscount        ForcedFloat            `json:"total_discount,omitempty"`
	TotalTax             ForcedFloat            `json:"total_tax,omitempty"`
	TotalPriceSet        *MultiCurrencySet      `json:"total_price_set,omitempty"`
	TaxLines             []TaxLine              `json:"tax_lines,omitempty"`
	PresentmentCurrency  string                 `json:"presentment_currency,omitempty"`
	FinancialStatus      FinancialStatus        `json:"financial_status,omitempty"`
	TaxIncluded          bool                   `json:"taxes_included,omitempty"`
	// These are only for creating orders
	SendReceipt            bool           `json:"send_receipt,omitempty"`
	SendFulfillmentReceipt bool           `json:"send_fulfillment_receipt,omitempty"`
	InventoryBehavior      string         `json:"inventory_behaviour,omitempty"` // Can be 'bypass', 'decrement_ignoring_policy', or 'decrement_obeying_policy'
	ShippingLines          []ShippingLine `json:"shipping_lines,omitempty"`
	// BuyXGetYLineItemID is used to keep track of line item ids, quantity which are part of BuyXGetY discounts.
	BuyXGetYLineItemIDToQty map[int64]int `json:"-"`
	// XXX more Shopify attributes
}

type ShippingLine struct {
	Code  string `json:"code"`
	Price string `json:"price"`
	Title string `json:"title"`
}

func HasTag(tags, searchTag string) bool {
	for _, tag := range strings.Split(tags, ",") {
		if strings.TrimSpace(tag) == searchTag {
			return true
		}
	}
	return false
}

type FinancialStatus string

type ExchangeOrders struct {
	OrderNo              int64                 `json:"order_no,omitempty"`
	OrdersExchangeHist   []*ExchangeOrdersHist `json:"exchange_orders_hist,omitempty"`
	NumberExchange       int64                 `json:"exchange_number,omitempty"`
	ExchangeLimitReached bool                  `json:"exchange_limit_reached,omitempty"`
}

type ExchangeOrdersHist struct {
	OrderID       int64 `json:"order_id,omitempty"`
	ParentOrderID int64 `json:"parent_order_id,omitempty"`
}

type Properties map[string]interface{}

func (p Properties) MarshalJSON() ([]byte, error) {
	out := []map[string]interface{}{}
	for k, v := range p {
		out = append(out, map[string]interface{}{
			"name":  k,
			"value": v,
		})
	}
	return json.Marshal(&out)
}

func (p *Properties) UnmarshalJSON(b []byte) error {
	in := []map[string]interface{}{}
	if err := json.Unmarshal(b, &in); err != nil {
		return err
	}
	(*p) = Properties{}
	for _, row := range in {
		(*p)[row["name"].(string)] = row["value"]
	}
	return nil
}

type DiscountAllocation struct {
	Amount                   string `json:"amount,omitempty"`
	DiscountApplicationIndex *int   `json:"discount_application_index,omitempty"`
}

type OrderItem struct {
	ID                  int64                 `json:"id,omitempty"`
	Name                string                `json:"name,omitempty"`
	Price               string                `json:"price,omitempty" description:"The price of the item before discounts have been applied in the shop currency."`
	ProductID           int64                 `json:"product_id,omitempty"`
	Quantity            int                   `json:"quantity,omitempty"`
	TaxLines            []TaxLine             `json:"tax_lines,omitempty"`
	Title               string                `json:"title,omitempty"`
	TotalDiscount       string                `json:"total_discount,omitempty"`
	VariantID           int64                 `json:"variant_id,omitempty"`
	VariantTitle        string                `json:"variant_title,omitempty"`
	Vendor              string                `json:"vendor,omitempty"`
	SKU                 string                `json:"sku,omitempty"`
	PreTaxPrice         string                `json:"pre_tax_price,omitempty"`
	IsGiftCard          bool                  `json:"gift_card,omitempty"`
	Properties          Properties            `json:"properties,omitempty"`
	PriceSet            *MultiCurrencySet     `json:"price_set,omitempty"`
	DiscountAllocations []*DiscountAllocation `json:"discount_allocations,omitempty"`
	// XXX more Shopify attributes
}

type TaxLine struct {
	Price ForcedFloat `json:"price"`
	Rate  float64     `json:"rate"`
	Title string      `json:"title"`
}

type Customer struct {
	ID             int64              `json:"id"`
	Email          string             `json:"email"`
	FirstName      string             `json:"first_name"`
	LastName       string             `json:"last_name"`
	DefaultAddress *ShippingAddress   `json:"default_address"`
	Addresses      []*ShippingAddress `json:"addresses"`
	Tags           string             `json:"tags"`
	VerifiedEmail  bool               `json:"verified_email"`
	// XXX more Shopify attributes
}

type ShippingAddress struct {
	Address1     string  `json:"address1,omitempty"`
	Address2     string  `json:"address2,omitempty"`
	City         string  `json:"city,omitempty"`
	Company      string  `json:"company,omitempty"`
	Country      string  `json:"country,omitempty"`
	FirstName    string  `json:"first_name,omitempty"`
	LastName     string  `json:"last_name,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	Province     string  `json:"province,omitempty"`
	Zip          string  `json:"zip,omitempty"`
	Name         string  `json:"name,omitempty"`
	CountryCode  string  `json:"country_code,omitempty"`  // BAD JSON FORMAT, should be countryCode
	ProvinceCode string  `json:"province_code,omitempty"` // BAD JSON FORMAT, should be provinceCode
}

type AppliedDiscount struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ValueType   string `json:"value_type,omitempty"`
	Value       string `json:"value,omitempty"`
	Amount      string `json:"amount,omitempty"`
}

// Corresponds to table return_shopping_draft_orders row
type ReturnShoppingDraft struct {
	ID              string     `json:"id,omitempty"`
	RetailerID      string     `json:"retailer_id,omitempty"`
	DraftOrderID    string     `json:"draft_order_id,omitempty"`
	OrderID         string     `json:"order_id,omitempty"`
	UpsellOrderID   string     `json:"upsell_order_id,omitempty"`
	ReturnSessionID string     `json:"return_session_id,omitempty"`
	Currency        string     `json:"currency,omitempty"`
	TotalAmount     float64    `json:"total_amount,omitempty"`
	TotalAmountPaid float64    `json:"total_amount_paid,omitempty"`
	InvoiceURL      string     `json:"invoice_url,omitempty"`
	PaidAt          *time.Time `json:"paid_at,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

type ShopifyDiscountRedistribution struct {
	ID                    string
	OrderID               string
	ReturnID              string
	RetailerID            string
	LineItemID            string
	Quantity              int
	Price                 float64
	OriginalDiscount      float64
	RedistributedDiscount float64
	CreatedAt             *time.Time
}

type InventoryLevel struct {
	InventoryItemID int64  `json:"inventory_item_id"`
	LocationID      int64  `json:"location_id"`
	Available       int64  `json:"available"`
	UpdatedAt       string `json:"updated_at"`
}

type Fulfillment struct {
	ID         int64       `json:"id"`
	OrderID    int64       `json:"order_id"`
	CreatedAt  string      `json:"created_at"`
	UpdatedAt  string      `json:"updated_at"`
	LineItems  []OrderItem `json:"line_items"`
	Status     string      `json:"status"`
	LocationID int64       `json:"location_id"`
}

type Product struct {
	ID             int64            `json:"id"`
	Title          string           `json:"title"`
	BodyHTML       string           `json:"body_html"`
	CreatedAt      string           `json:"created_at"`
	Vendor         string           `json:"vendor"`
	Tags           string           `json:"tags"`
	PublishedScope string           `json:"published_scope"`
	Images         []ProductImage   `json:"images"`
	Variants       []ProductVariant `json:"variants"`
	Options        []ProductOption  `json:"options"`
}

type ProductOption struct {
	ID        int64    `json:"id"`
	ProductID int64    `json:"product_id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Values    []string `json:"values"`
}

type ProductVariant struct {
	ID                int64   `json:"id"`
	Title             string  `json:"title"`
	Barcode           string  `json:"barcode"`
	ProductID         int64   `json:"product_id"`
	ImageID           int64   `json:"image_id"`
	SKU               string  `json:"sku"`
	Weight            float64 `json:"weight"`
	WeightUnit        string  `json:"weight_unit"`
	Price             string  `json:"price"`
	Position          int     `json:"position"`
	InventoryQuantity int64   `json:"inventory_quantity"`
	InventoryItemID   int64   `json:"inventory_item_id"`
	InventoryPolicy   string  `json:"inventory_policy"`
	// These are weird, as they link up with the product options
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
	Option3 string `json:"option3"`
	// Option4 doesn't exist on Shopify's ProductVariant but we use it when the
	// exchange by collection attribute adds the 4th option
	Option4 string `json:"option4"`
}

type ProductImage struct {
	CreatedAt  string  `json:"created_at"`
	ID         int64   `json:"id"`
	Position   int     `json:"position"`
	ProductID  int64   `json:"product_id"`
	VariantIDs []int64 `json:"variant_ids"`
	URL        string  `json:"src"`
	UpdatedAt  string  `json:"updated_at"`
}

type GiftCard struct {
	ID           int64       `json:"id,omitempty"`
	UserID       int64       `json:"user_id,omitempty"`
	OrderID      int64       `json:"order_id,omitempty"`
	CustomerID   int64       `json:"customer_id,omitempty"`
	LineItemID   int64       `json:"line_item_id,omitempty"`
	APIClientID  int64       `json:"api_client_id,omitempty"`
	InitialValue ForcedFloat `json:"initial_value,omitempty"` // need to force this to float
	Balance      ForcedFloat `json:"balance,omitempty"`       // need to force this to float
	Currency     string      `json:"currency,omitempty"`
	Note         string      `json:"note,omitempty"`
	CreatedAt    string      `json:"created_at,omitempty"`
	UpdatedAt    string      `json:"updated_at,omitempty"`
	DisabledAt   string      `json:"disabled_at,omitempty"`
	ExpiresOn    string      `json:"expires_on,omitempty"`
}

type Location struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Address1         string `json:"address1,omitempty"`
	Address2         string `json:"address2,omitempty"`
	City             string `json:"city,omitempty"`
	Zip              string `json:"zip,omitempty"`
	Province         string `json:"province,omitempty"`
	Country          string `json:"country,omitempty"`
	Phone            string `json:"phone,omitempty"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	CountryCode      string `json:"country_code,omitempty"`
	CountryName      string `json:"country_name,omitempty"`
	ProvinceCode     string `json:"province_code,omitempty"`
	Legacy           bool   `json:"legacy,omitempty"`
	AdmainGraphqlAPI string `json:"admin_graphql_api_id"`
}

type InventoryItem struct {
	Cost                 ForcedFloat `json:"cost"`
	CountryCodeOfOrigin  string      `json:"country_code_of_origin,omitempty"`
	HarmonizedCode       string      `json:"harmonized_system_code,omitempty"`
	ID                   int64       `json:"id"`
	ProvinceCodeOfOrigin string      `json:"province_code_of_origin,omitempty"`
	SKU                  string      `json:"sku,omitempty"`
	Tracked              bool        `json:"tracked"`
	RequiresShipping     bool        `json:"requires_shipping"`
}

type ForcedFloat float64
type ForcedString string

// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package domains

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateDomainRequest struct {
	Domain Domain `json:"domain"`
	// Client-supplied resource ID for the new domain. If omitted, the server
	// generates one.
	DomainId string `json:"-" url:"domain_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateDomainRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDomainRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDomainRequest struct {
	// When false (default), DeleteDomain is rejected with FAILED_PRECONDITION
	// if the domain still has Glossary pages. When true, those pages are
	// deleted first and then the domain is removed.
	Force bool `json:"-" url:"force,omitempty"`
	// Full resource name of the domain to delete. Format: `domains/{domain_id}`
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDomainRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDomainRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Domain struct {
	// Principal IDs of the business owners (users, groups, or service
	// principals).
	BusinessOwnerIds []int64 `json:"business_owner_ids,omitempty"`
	// Timestamp when the domain was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Full description (max 4096 chars)
	Description string `json:"description,omitempty"`
	// Unique identifier for the domain. If omitted at Create, the server
	// generates one.
	DomainId string `json:"domain_id,omitempty"`
	// Whether to mark the domain as a draft. If omitted on Create, the server
	// applies a default; the resolved value is returned in `effective_draft`.
	Draft bool `json:"draft,omitempty"`
	// Resolved draft state of the domain.
	EffectiveDraft bool `json:"effective_draft,omitempty"`
	// Icon to display for the domain.
	Icon *DomainIcon `json:"icon,omitempty"`
	// Full resource name of the domain. The primary identifier for this
	// resource. Format: `domains/{domain_id}` Identifies the domain on get,
	// update, and delete. Not an input on create — to choose the id, set
	// `CreateDomainRequest.domain_id`.
	Name string `json:"name,omitempty"`
	// Domain ID of the parent. If absent, this is a top-level domain. If
	// present, this domain is a subdomain of the specified parent.
	ParentDomainId string `json:"parent_domain_id,omitempty"`
	// Short description (max 280 chars)
	Subtitle string `json:"subtitle,omitempty"`
	// Governed tag key associated with this domain.
	TagKey string `json:"tag_key"`
	// Principal IDs of the technical owners (users, groups, or service
	// principals).
	TechnicalOwnerIds []int64 `json:"technical_owner_ids,omitempty"`
	// Timestamp when the domain was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Domain) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Domain) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Icon configuration for a domain.
type DomainIcon struct {
	// Hex color code with # prefix (e.g., "#FF5733").
	Color string `json:"color,omitempty"`

	Name DomainIconName `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DomainIcon) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DomainIcon) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Available icon names for a domain.
type DomainIconName string

const DomainIconNameAddressBook DomainIconName = `ADDRESS_BOOK`

const DomainIconNameAlarm DomainIconName = `ALARM`

const DomainIconNameArrowsIn DomainIconName = `ARROWS_IN`

const DomainIconNameAtom DomainIconName = `ATOM`

const DomainIconNameBalloon DomainIconName = `BALLOON`

const DomainIconNameBank DomainIconName = `BANK`

const DomainIconNameBarricade DomainIconName = `BARRICADE`

const DomainIconNameBasket DomainIconName = `BASKET`

const DomainIconNameBridge DomainIconName = `BRIDGE`

const DomainIconNameCactus DomainIconName = `CACTUS`

const DomainIconNameCallBell DomainIconName = `CALL_BELL`

const DomainIconNameCarrot DomainIconName = `CARROT`

const DomainIconNameChartPieSlice DomainIconName = `CHART_PIE_SLICE`

const DomainIconNameCity DomainIconName = `CITY`

const DomainIconNameCloud DomainIconName = `CLOUD`

const DomainIconNameCoins DomainIconName = `COINS`

const DomainIconNameCompassRose DomainIconName = `COMPASS_ROSE`

const DomainIconNameCraneTower DomainIconName = `CRANE_TOWER`

const DomainIconNameCrown DomainIconName = `CROWN`

const DomainIconNameCubeTransparent DomainIconName = `CUBE_TRANSPARENT`

const DomainIconNameFaders DomainIconName = `FADERS`

const DomainIconNameFlagBannerFold DomainIconName = `FLAG_BANNER_FOLD`

const DomainIconNameFlagCheckered DomainIconName = `FLAG_CHECKERED`

const DomainIconNameGavel DomainIconName = `GAVEL`

const DomainIconNameHamburger DomainIconName = `HAMBURGER`

const DomainIconNameHeadCircuit DomainIconName = `HEAD_CIRCUIT`

const DomainIconNameHourglassHigh DomainIconName = `HOURGLASS_HIGH`

const DomainIconNameIntersectThree DomainIconName = `INTERSECT_THREE`

const DomainIconNameMicroscope DomainIconName = `MICROSCOPE`

const DomainIconNameMoonStars DomainIconName = `MOON_STARS`

const DomainIconNamePackage DomainIconName = `PACKAGE`

const DomainIconNameParachute DomainIconName = `PARACHUTE`

const DomainIconNamePepper DomainIconName = `PEPPER`

const DomainIconNamePiggyBank DomainIconName = `PIGGY_BANK`

const DomainIconNamePill DomainIconName = `PILL`

const DomainIconNamePlanet DomainIconName = `PLANET`

const DomainIconNamePlant DomainIconName = `PLANT`

const DomainIconNamePlugsConnected DomainIconName = `PLUGS_CONNECTED`

const DomainIconNamePopcorn DomainIconName = `POPCORN`

const DomainIconNamePresentationChart DomainIconName = `PRESENTATION_CHART`

const DomainIconNamePuzzlePiece DomainIconName = `PUZZLE_PIECE`

const DomainIconNameRainbow DomainIconName = `RAINBOW`

const DomainIconNameRanking DomainIconName = `RANKING`

const DomainIconNameReceipt DomainIconName = `RECEIPT`

const DomainIconNameRocket DomainIconName = `ROCKET`

const DomainIconNameRuler DomainIconName = `RULER`

const DomainIconNameSailboat DomainIconName = `SAILBOAT`

const DomainIconNameScales DomainIconName = `SCALES`

const DomainIconNameScanSmiley DomainIconName = `SCAN_SMILEY`

const DomainIconNameScroll DomainIconName = `SCROLL`

const DomainIconNameShieldCheckered DomainIconName = `SHIELD_CHECKERED`

const DomainIconNameSneaker DomainIconName = `SNEAKER`

const DomainIconNameSnowflake DomainIconName = `SNOWFLAKE`

const DomainIconNameSolarRoof DomainIconName = `SOLAR_ROOF`

const DomainIconNameSpeedometer DomainIconName = `SPEEDOMETER`

const DomainIconNameStamp DomainIconName = `STAMP`

const DomainIconNameSteps DomainIconName = `STEPS`

const DomainIconNameStrategy DomainIconName = `STRATEGY`

const DomainIconNameSword DomainIconName = `SWORD`

const DomainIconNameTelevisionSimple DomainIconName = `TELEVISION_SIMPLE`

const DomainIconNameTent DomainIconName = `TENT`

const DomainIconNameTicket DomainIconName = `TICKET`

const DomainIconNameTractor DomainIconName = `TRACTOR`

const DomainIconNameTrafficCone DomainIconName = `TRAFFIC_CONE`

const DomainIconNameTrain DomainIconName = `TRAIN`

const DomainIconNameTreeEvergreen DomainIconName = `TREE_EVERGREEN`

const DomainIconNameTreeStructure DomainIconName = `TREE_STRUCTURE`

const DomainIconNameTrolleySuitcase DomainIconName = `TROLLEY_SUITCASE`

const DomainIconNameTrophy DomainIconName = `TROPHY`

const DomainIconNameTruckTrailer DomainIconName = `TRUCK_TRAILER`

const DomainIconNameUsersThree DomainIconName = `USERS_THREE`

const DomainIconNameVectorThree DomainIconName = `VECTOR_THREE`

// String representation for [fmt.Print]
func (f *DomainIconName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DomainIconName) Set(v string) error {
	switch v {
	case `ADDRESS_BOOK`, `ALARM`, `ARROWS_IN`, `ATOM`, `BALLOON`, `BANK`, `BARRICADE`, `BASKET`, `BRIDGE`, `CACTUS`, `CALL_BELL`, `CARROT`, `CHART_PIE_SLICE`, `CITY`, `CLOUD`, `COINS`, `COMPASS_ROSE`, `CRANE_TOWER`, `CROWN`, `CUBE_TRANSPARENT`, `FADERS`, `FLAG_BANNER_FOLD`, `FLAG_CHECKERED`, `GAVEL`, `HAMBURGER`, `HEAD_CIRCUIT`, `HOURGLASS_HIGH`, `INTERSECT_THREE`, `MICROSCOPE`, `MOON_STARS`, `PACKAGE`, `PARACHUTE`, `PEPPER`, `PIGGY_BANK`, `PILL`, `PLANET`, `PLANT`, `PLUGS_CONNECTED`, `POPCORN`, `PRESENTATION_CHART`, `PUZZLE_PIECE`, `RAINBOW`, `RANKING`, `RECEIPT`, `ROCKET`, `RULER`, `SAILBOAT`, `SCALES`, `SCAN_SMILEY`, `SCROLL`, `SHIELD_CHECKERED`, `SNEAKER`, `SNOWFLAKE`, `SOLAR_ROOF`, `SPEEDOMETER`, `STAMP`, `STEPS`, `STRATEGY`, `SWORD`, `TELEVISION_SIMPLE`, `TENT`, `TICKET`, `TRACTOR`, `TRAFFIC_CONE`, `TRAIN`, `TREE_EVERGREEN`, `TREE_STRUCTURE`, `TROLLEY_SUITCASE`, `TROPHY`, `TRUCK_TRAILER`, `USERS_THREE`, `VECTOR_THREE`:
		*f = DomainIconName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADDRESS_BOOK", "ALARM", "ARROWS_IN", "ATOM", "BALLOON", "BANK", "BARRICADE", "BASKET", "BRIDGE", "CACTUS", "CALL_BELL", "CARROT", "CHART_PIE_SLICE", "CITY", "CLOUD", "COINS", "COMPASS_ROSE", "CRANE_TOWER", "CROWN", "CUBE_TRANSPARENT", "FADERS", "FLAG_BANNER_FOLD", "FLAG_CHECKERED", "GAVEL", "HAMBURGER", "HEAD_CIRCUIT", "HOURGLASS_HIGH", "INTERSECT_THREE", "MICROSCOPE", "MOON_STARS", "PACKAGE", "PARACHUTE", "PEPPER", "PIGGY_BANK", "PILL", "PLANET", "PLANT", "PLUGS_CONNECTED", "POPCORN", "PRESENTATION_CHART", "PUZZLE_PIECE", "RAINBOW", "RANKING", "RECEIPT", "ROCKET", "RULER", "SAILBOAT", "SCALES", "SCAN_SMILEY", "SCROLL", "SHIELD_CHECKERED", "SNEAKER", "SNOWFLAKE", "SOLAR_ROOF", "SPEEDOMETER", "STAMP", "STEPS", "STRATEGY", "SWORD", "TELEVISION_SIMPLE", "TENT", "TICKET", "TRACTOR", "TRAFFIC_CONE", "TRAIN", "TREE_EVERGREEN", "TREE_STRUCTURE", "TROLLEY_SUITCASE", "TROPHY", "TRUCK_TRAILER", "USERS_THREE", "VECTOR_THREE"`, v)
	}
}

// Values returns all possible values for DomainIconName.
//
// There is no guarantee on the order of the values in the slice.
func (f *DomainIconName) Values() []DomainIconName {
	return []DomainIconName{
		DomainIconNameAddressBook,
		DomainIconNameAlarm,
		DomainIconNameArrowsIn,
		DomainIconNameAtom,
		DomainIconNameBalloon,
		DomainIconNameBank,
		DomainIconNameBarricade,
		DomainIconNameBasket,
		DomainIconNameBridge,
		DomainIconNameCactus,
		DomainIconNameCallBell,
		DomainIconNameCarrot,
		DomainIconNameChartPieSlice,
		DomainIconNameCity,
		DomainIconNameCloud,
		DomainIconNameCoins,
		DomainIconNameCompassRose,
		DomainIconNameCraneTower,
		DomainIconNameCrown,
		DomainIconNameCubeTransparent,
		DomainIconNameFaders,
		DomainIconNameFlagBannerFold,
		DomainIconNameFlagCheckered,
		DomainIconNameGavel,
		DomainIconNameHamburger,
		DomainIconNameHeadCircuit,
		DomainIconNameHourglassHigh,
		DomainIconNameIntersectThree,
		DomainIconNameMicroscope,
		DomainIconNameMoonStars,
		DomainIconNamePackage,
		DomainIconNameParachute,
		DomainIconNamePepper,
		DomainIconNamePiggyBank,
		DomainIconNamePill,
		DomainIconNamePlanet,
		DomainIconNamePlant,
		DomainIconNamePlugsConnected,
		DomainIconNamePopcorn,
		DomainIconNamePresentationChart,
		DomainIconNamePuzzlePiece,
		DomainIconNameRainbow,
		DomainIconNameRanking,
		DomainIconNameReceipt,
		DomainIconNameRocket,
		DomainIconNameRuler,
		DomainIconNameSailboat,
		DomainIconNameScales,
		DomainIconNameScanSmiley,
		DomainIconNameScroll,
		DomainIconNameShieldCheckered,
		DomainIconNameSneaker,
		DomainIconNameSnowflake,
		DomainIconNameSolarRoof,
		DomainIconNameSpeedometer,
		DomainIconNameStamp,
		DomainIconNameSteps,
		DomainIconNameStrategy,
		DomainIconNameSword,
		DomainIconNameTelevisionSimple,
		DomainIconNameTent,
		DomainIconNameTicket,
		DomainIconNameTractor,
		DomainIconNameTrafficCone,
		DomainIconNameTrain,
		DomainIconNameTreeEvergreen,
		DomainIconNameTreeStructure,
		DomainIconNameTrolleySuitcase,
		DomainIconNameTrophy,
		DomainIconNameTruckTrailer,
		DomainIconNameUsersThree,
		DomainIconNameVectorThree,
	}
}

// Type always returns DomainIconName to satisfy [pflag.Value] interface
func (f *DomainIconName) Type() string {
	return "DomainIconName"
}

type GetDomainRequest struct {
	// Full resource name of the domain to retrieve. Format:
	// `domains/{domain_id}`
	Name string `json:"-" url:"-"`
}

func (s *GetDomainRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListDomainsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Filter by parent domain. - Absent: return all domains regardless of
	// hierarchy. - Present: return only direct children of the specified
	// domain.
	ParentDomainId string `json:"-" url:"parent_domain_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDomainsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDomainsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDomainsResponse struct {
	Domains []Domain `json:"domains,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDomainsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDomainsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateDomainRequest struct {
	Domain Domain `json:"domain"`
	// Full resource name of the domain. The primary identifier for this
	// resource. Format: `domains/{domain_id}` Identifies the domain on get,
	// update, and delete. Not an input on create — to choose the id, set
	// `CreateDomainRequest.domain_id`.
	Name string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateDomainRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

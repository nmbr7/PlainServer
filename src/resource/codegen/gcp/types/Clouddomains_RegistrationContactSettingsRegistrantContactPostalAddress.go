package types

type Clouddomains_RegistrationContactSettingsRegistrantContactPostalAddress struct {
	/*
	   Required. CLDR region code of the country/region of the address. This is never inferred and it is up to the user to
	   ensure the value is correct. See https://cldr.unicode.org/ and
	   https://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: "CH" for Switzerland.
	*/
	RegionCode string `json:"regionCode,omitempty" yaml:"regionCode,omitempty"`

	/*
	   Unstructured address lines describing the lower levels of an address.
	   Because values in addressLines do not have type information and may sometimes contain multiple values in a single
	   field (e.g. "Austin, TX"), it is important that the line order is clear. The order of address lines should be
	   "envelope order" for the country/region of the address. In places where this can vary (e.g. Japan), address_language
	   is used to make it explicit (e.g. "ja" for large-to-small ordering and "ja-Latn" or "en" for small-to-large). This way,
	   the most specific line of an address can be selected based on the language.
	*/
	AddressLines []string `json:"addressLines,omitempty" yaml:"addressLines,omitempty"`

	/*
	   Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state,
	   a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community
	   (e.g. "Barcelona" and not "Catalonia"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland
	   this should be left unpopulated.
	*/
	AdministrativeArea string `json:"administrativeArea,omitempty" yaml:"administrativeArea,omitempty"`

	/*
	   Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world
	   where localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.
	*/
	Locality string `json:"locality,omitempty" yaml:"locality,omitempty"`

	// The name of the organization at the address.
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`

	/*
	   Postal code of the address. Not all countries use or require postal codes to be present, but where they are used,
	   they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
	*/
	PostalCode string `json:"postalCode,omitempty" yaml:"postalCode,omitempty"`

	/*
	   The recipient at the address. This field may, under certain circumstances, contain multiline information. For example,
	   it might contain "care of" information.

	   - - -
	*/
	Recipients []string `json:"recipients,omitempty" yaml:"recipients,omitempty"`
}

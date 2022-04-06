package main

import (
  "github.com/alvinbaena/passkit"
)

func main() {
c := passkit.NewBoardingPass(passkit.TransitTypeAir)
field := passkit.Field{
    Key: "key",
    Label: "label",
    Value:"value",
}

c.AddHeaderField(field)
c.AddPrimaryFields(field)
c.AddSecondaryFields(field)

p := passkit.Pass{
    FormatVersion:       1,
    TeamIdentifier:      "TEAMID",
    PassTypeIdentifier:  "pass.type.id",
    AuthenticationToken: "123141lkjdasj12314",
    OrganizationName:    "Your Organization",
    SerialNumber:        "1234",
    Description:         "test",
    BoardingPass:         c,
    Barcodes: []passkit.Barcode{
        {
            Format:          passkit.BarcodeFormatPDF417,
            Message:         "1312312312312312312312312312",
            MessageEncoding: "utf-8",
        },
    },
}
}

package entities

type ContentTransferEncoding string

const (
	QuotedPrintable ContentTransferEncoding = "quoted-printable\u000d"
	The7Bit         ContentTransferEncoding = "7bit\u000d"
)

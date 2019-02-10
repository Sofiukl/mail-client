package cmd

/*Notification structure
* for primarily Email
 */
type Notification struct {
	To      []string
	Subject string
	Message string
}

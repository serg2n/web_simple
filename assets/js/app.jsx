class ContactItem extends React.Component {
    render() {
        return (
            <tr>
                <td> {this.props.Id}    </td>
                <td> {this.props.FirstName} </td>
                <td> {this.props.LastName}  </td>
                <td> {this.props.Phone}  </td>
                <td> {this.props.Email}  </td>
            </tr>
        );
    }
}

class ContactList extends React.Component {
    constructor(props) {
        super(props);
        this.state = { contacts: [] };
    }

    componentDidMount() {
        this.serverRequest =
            axios
                .get("/contact")
                .then((result) => {
                    this.setState({ contacts: result.data });
                });
    }

    render() {
        const contacts = this.state.contacts.map((contact, i) => {
            return (
                <ContactItem key={i} Id={contact.Id} FirstName={contact.FirstName} LastName={contact.LastName} Phone={contact.Phone} Email={contact.Email}/>
            );
        });

        return (
            <div>
                <table><tbody>
                <tr><th>Id</th><th>First Name</th><th>Last Name</th><th>Phone</th><th>Email</th></tr>
                {contacts}
                </tbody></table>

            </div>
        );
    }
}

ReactDOM.render(<ContactList/>, document.querySelector("#root"));
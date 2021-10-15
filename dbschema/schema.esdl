
module Meetup {

    type User {
        property username -> str;
        property email -> str {
            delegated constraint exclusive
        }
    };

    type Meetup {
        property name -> str;
        property description -> str;
        property userId -> str
    };

}

CREATE MIGRATION m1272aqowfivchn7errl4cojtmslhmaqhecjiqnc2ccmvpcnpyddqq
    ONTO initial
{
  CREATE MODULE Meetup IF NOT EXISTS;
  CREATE TYPE Meetup::Meetup {
      CREATE PROPERTY description -> std::str;
      CREATE PROPERTY name -> std::str;
      CREATE PROPERTY userId -> std::uuid;
  };
  CREATE TYPE Meetup::User {
      CREATE PROPERTY email -> std::str {
          CREATE DELEGATED CONSTRAINT std::exclusive;
      };
      CREATE PROPERTY username -> std::str;
  };
};

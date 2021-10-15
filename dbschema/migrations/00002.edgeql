CREATE MIGRATION m1eqocsdmgd7mv2cv4j3j2bgbwyxaoo3wvu4764x66z3anh3sjdsfa
    ONTO m1272aqowfivchn7errl4cojtmslhmaqhecjiqnc2ccmvpcnpyddqq
{
  ALTER TYPE Meetup::Meetup {
      ALTER PROPERTY userId {
          SET TYPE std::str USING ('');
      };
  };
};

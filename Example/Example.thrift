namespace swift TRE

include "Shared.thrift"

struct TFoo {}

struct TExample {
    1:required double amountAvailable;
    2:required string rebateAmountAvailable;
    3:optional bool amountPendingVerification;
    4:optional i32 pendingWithdrawAmount;
    5:optional TServiceType serviceType;
    5:optional i64 unpaidAmount;
    6:required TFoo fooe;
    7:required list<TFoo> fooes;
    8:required list<string> strs;
    9:required list<i16> ints;
    10:required list<Shared.TBasic> basics;
    11:required list<i64> int64s;
}

service Example extends Shared.Basic {
  void ping(1: string ip),
  i32 getPendingWithdrawAmount()
  list<TServiceType> getAvailableServieType(1: string id, 2: TFoo fooe, 3: list<string> names)
}

enum TServiceType {
  Buy4Me = 1
  Ship4Me = 2
  Ezbuy = 3
  Prime = 4
}
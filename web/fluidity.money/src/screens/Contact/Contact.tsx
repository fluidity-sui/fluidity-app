// Copyright 2022 Fluidity Money. All rights reserved. Use of this
// source code is governed by a GPL-style license that can be found in the
// LICENSE.md file.

import { GeneralButton } from "@fluidity-money/surfing";
import styles from "./Contact.module.scss";

const Contact = () => {
  return (
    <div className={styles.container}>
      <div>Contact</div>
      <GeneralButton
        type={"primary"}
        size={"small"}
        handleClick={function (): void {
          throw new Error("Function not implemented.");
        }}
      >
        SUBSCRIBE
      </GeneralButton>
    </div>
  );
};

export default Contact;

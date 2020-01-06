package controller

import (
    "bufio"
)

func write(writer bufio.Writer, str string) error {
    if _, err := writer.WriteString(str); err != nil {
        return err
    }

    if err := writer.Flush(); err != nil {
        return err
    }

    return nil
}

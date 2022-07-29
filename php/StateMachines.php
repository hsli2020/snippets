<?php # https://github.com/JacobBennett/StateMachinesTalk

# app/Enums/InvoiceState.php
namespace App\Enums;

enum InvoiceState: string
{
    case Draft = 'draft';
    case Open = 'open';
    case Paid = 'paid';
    case Void = 'void';
    case Uncollectable = 'uncollectable';
}

# app/Http/Controllers/CancelInvoiceController.php
namespace App\Http\Controllers;

use App\Models\Invoice;

class CancelInvoiceController extends Controller
{
    public function __invoke(Request $request, Invoice $invoice)
    {
        $invoice->state()->cancel();
        // $invoice->state()->finalize();
        // $invoice->update($request->validate());
        // $invoice->state()->pay();
        // $invoice->state()->void();
        return view('invoice.show', ['invoice' => $invoice]);
    }
}

# app/Models/Invoice.php

namespace App\Models;

use App\Enums\InvoiceState;
use http\Exception\InvalidArgumentException;
use App\StateMachines\Invoice\PaidInvoiceState;
use App\StateMachines\Invoice\VoidInvoiceState;
use App\StateMachines\Invoice\OpenInvoiceState;
use App\StateMachines\Invoice\DraftInvoiceState;
use App\StateMachines\Invoice\InvoiceStateContract;
use App\StateMachines\Invoice\UncollectableInvoiceState;

class Invoice extends BaseModel
{
    protected $attributes = [
        'status' => InvoiceState::Draft,
    ];

    protected $casts = [
        'status' => InvoiceState::class,
    ];

    public function state(): InvoiceStateContract
    {
        return match ($this->status) {
            InvoiceState::Draft => new DraftInvoiceState($this),
            InvoiceState::Open => new OpenInvoiceState($this),
            InvoiceState::Paid => new PaidInvoiceState($this),
            InvoiceState::Void => new VoidInvoiceState($this),
            InvoiceState::Uncollectable => new UncollectableInvoiceState($this),
            default => throw new InvalidArgumentException('Invalid status'),
        };
    }
}

# app/StateMachines/Invoice/InvoiceStateContract.php
namespace App\StateMachines\Invoice;

use App\Models\Invoice;
use App\Enums\InvoiceState;

interface InvoiceStateContract
{
    public function __construct(Invoice $invoice);
    public function finalize(): void;
    public function pay(): void;
    public function void(): void;
    public function cancel(): void;
}

# app/StateMachines/Invoice/BaseInvoiceState.php
namespace App\StateMachines\Invoice;

use App\Models\Invoice;

class BaseInvoiceState implements InvoiceStateContract
{
    function __construct(public Invoice $invoice) {}
    function finalize() { throw new \Exception(); }
    function pay() { throw new \Exception(); }
    function void() { throw new \Exception(); }
    function cancel() { throw new \Exception(); }
}

# app/StateMachines/Invoice/DraftInvoiceState.php
namespace App\StateMachines\Invoice;

use App\Enums\InvoiceState;

class DraftInvoiceState extends BaseInvoiceState
{
    function finalize() {
        $this->invoice->update(['status' => InvoiceState::Open]);
        /** Pseudo Code Below */
        Mail::send(new InvoiceDue($this->invoice));
    }
}

# app/StateMachines/Invoice/OpenInvoiceState.php
namespace App\StateMachines\Invoice;

use App\Enums\InvoiceState;

class OpenInvoiceState extends BaseInvoiceState
{
    function pay() {
        $this->invoice->update(['status' => InvoiceState::Paid]);
    }

    function void() {
        $this->invoice->update(['status' => InvoiceState::Void]);
    }

    function cancel() {
        $this->invoice->update(['status' => InvoiceState::Uncollectable]);
    }
}

# app/StateMachines/Invoice/UncollectableInvoiceState.php
namespace App\StateMachines\Invoice;

use App\Enums\InvoiceState;

class UncollectableInvoiceState extends BaseInvoiceState
{
    function pay() {
        $this->invoice->update(['status' => InvoiceState::Paid]);
        /* Pseudo Code Below */
        Mail::send(new Invoice($this->invoice));
    }

    function void() {
        $this->invoice->update(['status' => InvoiceState::Void]);
    }
}

# app/StateMachines/Invoice/PaidInvoiceState.php
namespace App\StateMachines\Invoice;

class PaidInvoiceState extends BaseInvoiceState
{
    /** FINAL STATE */
}

# app/StateMachines/Invoice/VoidInvoiceState.php
namespace App\StateMachines\Invoice;

class VoidInvoiceState extends BaseInvoiceState
{
    /** FINAL STATE */
}
